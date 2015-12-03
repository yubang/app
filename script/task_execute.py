# coding:UTF-8


"""
任务执行模块
@author: yubang
"""


from model.task_queue import TaskQueueModel
from lib.log import login_log


def init():
    tasks = TaskQueueModel.select().limit(1)
    for task in tasks:
        handle_task(task)


def add_one_container(task):
    """
    添加一个容器
    :return: boolean, 容器信息
    """
    return True, {}


def handle_task(task):
    if task.command_code == 1:
        # 添加一个容器
        r, _ = add_one_container(task)
        if r:
            login_log("common", "申请容器成功，触发应用id：%d"%task.app_id)
        else:
            login_log("error", "申请容器失败，触发应用id：%d"%task.app_id)
    elif task.command_code == 2:
        # 移除一个容器
        pass
    elif task.command_code == 3:
        # 移除所有容器
        pass
    else:
        # 未定义任务码
        login_log("error", "未定义任务码：%d，触发应用id：%d" % (task.command_code, task.app_id))