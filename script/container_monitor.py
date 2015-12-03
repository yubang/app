# coding:UTF-8


"""
监视应用是否需要移除或添加容器
@author: yubang
"""


from model.app import AppModel
from model.app_container import AppContainModel
from model.task_queue import TaskQueueModel
import datetime


def init():
    apps = AppModel.select()
    for app in apps:
        handle_app(app)


def handle_app(app):
    container_nums = AppContainModel.select().where(AppContainModel.app_id == app.id).count()

    # 判断是否需要增加容器
    if container_nums == 0:
        add_sign = 1
    else:
        add_sign = 0

    # 判断是否需要减少容器
    if container_nums > app.max_container_number:
        add_sign = 0
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
