# coding:UTF-8


"""
用户操作模块
@author: yubang
"""


from bottle import Bottle, static_file, request, abort, redirect
from model.app import AppModel
import os
import datetime
import json

user_app = Bottle()


def __output_html(html_name):
    """
    输出html文件
    :param html_name: 文件名（不要后缀）
    :return:
    """
    path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.dirname(path)
    path = os.path.join(path, "static/html/user")
    return static_file(html_name+".html", root=path)


def get_login_user_id():
    """
    获取user_id，没有登录返回0
    :return: int
    """
    return 1


def get_image_name_from_env(env):
    """
    根据索引获取镜像名称
    :param env: 索引
    :return: str
    """
    images = {'0': "python2.7", '1': "python3.5", "2": "go1.5", "3": "nodejs4", "4": "java8", "5": "static", "6": "php5.6"}
    return images.get(env, None)


@user_app.get("/")
def index():
    return __output_html("index")


@user_app.post("/addApp")
def add_app():

    user_id = get_login_user_id()

    title = request.forms.title
    description = request.forms.description
    min_number = request.forms.min_number
    max_number = request.forms.max_number
    memory = request.forms.memory
    env = request.forms.env
    app_host = request.forms.app_host
    app_port = request.forms.app_port
    code_address = request.forms.code_address

    if user_id == 0:
        return abort(403, '未登录！')

    env = get_image_name_from_env(env)
    if env is None:
        return abort(403, "env error!")

    dao_obj = AppModel(title=title, description=description, user_id=user_id, memory=memory, env=env,
                       code_address=code_address, app_host=app_host, app_port=app_port, min_container_number=min_number,
                       max_container_number=max_number, create_time=datetime.datetime.now())
    dao_obj.save()
    return redirect("/user")


@user_app.get("/getApps")
def get_apps():
    """
    获取用户所有应用
    :return: dict
    """
    user_id = get_login_user_id()
    apps = AppModel.select().where(AppModel.user_id==user_id)
    data = map(AppModel.get_dict_from_model_obj, apps)
    return json.dumps(list(data))


@user_app.get("/deleteApp/:app_id#\d+#")
def delete_app(app_id):
    """
    删除应用
    :param app_id 应用id
    :return:
    """
    user_id = get_login_user_id()
    dao = AppModel.delete().where(AppModel.user_id == user_id, AppModel.id == app_id)
    dao.execute()

    return redirect("/user")
