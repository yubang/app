# coding:UTF-8


"""
监视应用是否需要移除或添加容器
@author: yubang
"""


from model.app import AppModel
from model.app_container import AppContainModel
from model.task_queue import TaskQueueModel
from lib.config import get_config_data
import datetime
import requests
import json


def init():
    """
    入口函数
    :return:
    """
    apps = AppModel.select()
    for app in apps:
        handle_app(app)


def request_api(url, data):
    """
    请求容器API
    :param url: 容器API地址
    :param data: post参数
    :return: response对象
    """
    d = get_config_data()
    data['token'] = d['token.token']
    try:
        return requests.post(url, data)
    except:
        class Obj(object):
           status_code = 0
        obj = Obj()
        return obj


def get_app_avg_message(app_id):
    """
    获取应用平均状态
    :return: cpu，内存
    """
    cpu = 0
    memory = 0
    count = AppContainModel.select().where(AppContainModel.app_id == app_id).count()
    objs = AppContainModel.select().where(AppContainModel.app_id == app_id)
    for obj in objs:
        r = request_api(obj.api_url + "stats", {"containerId": obj.container_id})
        if r.status_code != 200:
            cpu += 100
            memory += 100
            continue
        d = json.loads(r.text)
        if d['code'] != 0:
            cpu += 100
            memory += 100
            continue

        cpu += d['result']['cpu']

        memory += d['result']['memory']

    if count == 0:
        return 100, 100

    cpu = cpu / count
    cpu = float("%.2f" % cpu)
    memory = memory / count
    memory = float("%.2f" % memory)

    return cpu, memory


def handle_app(app):
    """
    处理每一个APP
    :param app: app对象
    :return:
    """
    container_nums = AppContainModel.select().where(AppContainModel.app_id == app.id).count()

    cpu, memory = get_app_avg_message(app.id)

    # 判断是否需要增加容器
    if container_nums < app.min_container_number:
        add_sign = 1
    else:
        if cpu > 5.0 and memory > 95.5 and container_nums < app.max_container_number:
            add_sign = 1
        else:
            add_sign = 0

    # 判断是否需要减少容器
    if container_nums > app.max_container_number:
        add_sign = 0
        reduce_sign = -1
    else:
        if cpu < 1.5 and memory < 90 and container_nums > app.min_container_number:
            reduce_sign = -1
        else:
            reduce_sign = 0

    sign = add_sign + reduce_sign
    if sign == 0:
        return None

    obj = TaskQueueModel(user_id=0, app_id=app.id, create_time=datetime.datetime.now())
    obj.command_content = '{}'
    if sign == 1:
        obj.command_code = 1 # 增加容器
    elif sign == -1:
        obj.command_code = 2 # 减少容器

    # 判断命令有没有重复
    if TaskQueueModel.select().where(TaskQueueModel.user_id == obj.user_id, TaskQueueModel.app_id == obj.app_id, TaskQueueModel.command_code == obj.command_code).count():
        return None

    # 删除同应用不同操作
    dao = TaskQueueModel.delete().where(TaskQueueModel.user_id == obj.user_id, TaskQueueModel.app_id == obj.app_id)
    dao.execute()

    obj.save()
