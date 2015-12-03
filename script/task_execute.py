# coding:UTF-8


"""
任务执行模块
@author: yubang
"""


from model.task_queue import TaskQueueModel
from model.app_container import AppContainModel
from model.container_server import ContainerServerModel
from model.app import AppModel
from lib.log import login_log
from lib.config import get_config_data
import requests
import urllib
import json
import datetime
import os
import subprocess


def init():
    tasks = TaskQueueModel.select().limit(1)
    for task in tasks:
        handle_task(task)


def request_api(url, data):
    d = get_config_data()
    data['token'] = d['token.token']
    return requests.post(url, data)


def get_a_api_url():
    """
    获取一个API服务器地址
    :return:
    """
    try:
        obj = ContainerServerModel.select().where(ContainerServerModel.status == 0).get()
        url = "http://%s:%d/container/" % (obj.server_host, obj.server_port)
        return url, obj.server_host
    except:
        return None, None


def add_one_container(task):
    """
    添加一个容器
    :return: boolean, 容器信息
    """

    try:
        app_obj = AppModel.select().where(AppModel.id == task.app_id).get()
    except:
        return False, None

    api_url, api_host = get_a_api_url()
    if api_url is None:
        return False, {}

    d = api_url

    api_url += "build"
    data = {
        "code_address": urllib.request.quote(app_obj.code_address),
        "memory": app_obj.memory,
        "image_name": app_obj.env
    }
    r = request_api(api_url, data)

    if r.status_code != 200:
        return False, {}

    obj = json.loads(r.text)
    if obj['code'] != 0:
        return False, {}

    return True, {"r": obj, "host": api_host, "api_url": d}


def remove_some_container(task, number):
    """
    删除若干个容器
    :param task:
    :param number:
    :return:
    """

    containers = AppContainModel.select().where(AppContainModel.app_id == task.app_id).limit(number)
    for obj in containers:
        api_url = obj.api_url + "remove"
        data = {
            "containerId": obj.container_id
        }
        r = request_api(api_url, data)
        if r.status_code == 200:
            d = json.loads(r.text)
            if d['code'] == 0:
                # 删除容器成功
                dao = AppContainModel.delete().where(AppContainModel.id == obj.id)
                dao.execute()
                login_log("common", "删除容器成功，触发应用id：%d，容器id：%s，容器服务器域名：%s" % (task.app_id, obj.container_id, obj.host))
                continue
        # 删除出错处理
        login_log("error", "删除容器失败，触发应用id：%d，容器id：%s，容器服务器域名：%s" % (task.app_id, obj.container_id, obj.host))

    return None


def build_nginx_config(app_id):
    """
    生成应用nginx配置文件
    :param app_id:
    :return:
    """
    nginx_str = """

    upstream site%d {
        %s
    }

    server{
        listen %s:%d;
        index index.html index.htm;
        location / {
            proxy_set_header X-Forwarded-For $proxy_add_x_forwarded_for;
            proxy_set_header Host $http_host;
            proxy_pass http://site%d;
        }
    }
    """

    try:
        app = AppModel.select().where(AppModel.id == app_id).get()
    except:
        return False

    app_lists = ""
    containers = AppContainModel.select().where(AppContainModel.app_id == app_id)
    for obj in containers:
        app_lists = app_lists + ("server %s:%d;\n" % (obj.host, obj.port))

    config_path = os.path.dirname(os.path.dirname(os.path.realpath(__file__)))
    config_path = config_path + "/data/nginx_config/" + str(app_id) + ".conf"
    if os.path.exists(config_path):
        os.remove(config_path)

    if app_lists == "":
        return False

    config_str = nginx_str % (app_id, app_lists, app.app_host, app.app_port, app_id)
    fp = open(config_path, "w")
    fp.write(config_str)
    fp.close()

    subprocess.getstatusoutput("nginx -s reload")


def handle_task(task):
    if task.command_code == 1:
        # 添加一个容器
        r, obj = add_one_container(task)

        if r:
            # 保存容器服务器申请到的容器
            dao = AppContainModel(app_id=task.app_id, host=obj['host'], port=obj['r']['result']['port'],
                                  container_id=obj['r']['result']['containerId'], create_time=datetime.datetime.now(),
                                  api_url=obj['api_url'])
            dao.save()
            login_log("common", "申请容器成功，触发应用id：%d，申请到的容器域名：%s，端口：%d，id：%s" % (task.app_id,
                                                                              obj['host'], obj['r']['result']['port'],
                                                                              obj['r']['result']['containerId']))
        else:
            login_log("error", "申请容器失败，触发应用id：%d"%task.app_id)
    elif task.command_code == 2:
        # 移除一个容器
        remove_some_container(task, 1)
    elif task.command_code == 3:
        # 移除所有容器
        count = AppContainModel.select().where(AppContainModel.app_id == task.app_id).count()
        remove_some_container(task, count)
    else:
        # 未定义任务码
        login_log("error", "未定义任务码：%d，触发应用id：%d" % (task.command_code, task.app_id))

    build_nginx_config(task.app_id)

    # 删除任务队列
    dao = TaskQueueModel.delete().where(TaskQueueModel.id == task.id)
    dao.execute()

