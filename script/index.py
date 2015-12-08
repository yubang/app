# coding:UTF-8

"""
自动脚本
@author: yubang
"""


from script import container_monitor
from script import task_execute
from lib.log import login_log
from model.base import start_connect, close_connect
import threading
import time


class PassService(threading.Thread):
    def run(self):
        while True:

            # 打开连接
            start_connect()

            try:
                container_monitor.init()
                login_log("container_monitor_script", "完成一次应用状态扫描！")
            except Exception as err:
                login_log("container_monitor_script_error", err.args[0])
                time.sleep(5)

            try:
                if not task_execute.init():
                    login_log("task_script", "任务队列没有任务")
                login_log("task_script", "完成一次任务调度！")
            except Exception as err:
                login_log("task_script_error", err.args[0])
                time.sleep(5)

            # 关闭数据库
            close_connect()

            time.sleep(5)


def init():
    """
    初始化函数
    :return:
    """
    service = PassService()
    service.start()

    # container_monitor.init()
    # task_execute.init()