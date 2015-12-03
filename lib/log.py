# coding:UTF-8


"""
日记模块
@author: yubang
"""


import datetime
import os


def login_log(log_type, log_content):
    """
    记录日记
    :return:
    """

    data = "[%s] %s \n" % (datetime.datetime.now().strftime("%Y-%m-%d %H:%M:%S"), log_content)

    dir_path = os.path.dirname(os.path.dirname(os.path.realpath(__file__)))
    dir_path = os.path.join(dir_path, "data", "log")
    fp_path = os.path.join(dir_path, datetime.datetime.now().strftime("%Y%m%d")+"_"+log_type+".log")
    fp = open(fp_path, "a")
    fp.write(data)
    fp.close()
