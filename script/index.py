# coding:UTF-8

"""
自动脚本
@author: yubang
"""


from script import container_monitor
from script import task_execute

def init():
    """
    初始化函数
    :return:
    """
    print("init")

    container_monitor.init()
    task_execute.init()

    print("stop")

