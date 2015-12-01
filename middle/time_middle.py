# coding:UTF-8

"""
统计运行效率中间件
@author: yubang
"""


import datetime


class TimeMiddle(object):
    def before_request(self, request, response):
        self.__start_run_time = datetime.datetime.now().microsecond
        return None
    def after_request(self, request, response):
        self.__end_run_time = datetime.datetime.now().microsecond
        run_time = self.__end_run_time - self.__start_run_time
        print("runtime:" + str(run_time) + "ms")
        response.set_header("runtime", str(run_time) + "ms")
        return response