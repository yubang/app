# coding:UTF-8


"""
用户操作模块
@author: yubang
"""


from bottle import Bottle, static_file
import os


user_app = Bottle()


def __output_html(html_name):
    path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.dirname(path)
    path = os.path.join(path, "static/html/user")
    return static_file(html_name+".html", root=path)


@user_app.get("/")
def index():
    return  __output_html("index")
