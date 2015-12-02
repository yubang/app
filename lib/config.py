# coding:UTF-8


"""
加载项目配置文件
@author: yubang
"""


from bottle import default_app
import os


def get_config_data():
    config_file_path = os.path.dirname(os.path.dirname(os.path.realpath(__file__))) + "/data/config/account.conf"
    data = default_app().config.load_config(config_file_path)
    return data
