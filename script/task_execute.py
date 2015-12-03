# coding:UTF-8


"""
任务执行模块
@author: yubang
"""


from model.task_queue import TaskQueueModel


def init():
    tasks = TaskQueueModel.select().limit(1)
    for task in tasks:
        handle_task(task)


def handle_task(task):
    print(task)