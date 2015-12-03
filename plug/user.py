# coding:UTF-8


"""
用户操作模块
@author: yubang
"""


from bottle import Bottle, static_file, request, abort, redirect
from model.app import AppModel
from model.user import UserModel
from model.task_queue import TaskQueueModel
import os
import datetime
import json
import hashlib

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
    return request.session.get('uid', 0)


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


@user_app.post("/updateApp")
def update_app():
    """
    更新应用信息
    :return:
    """

    user_id = get_login_user_id()

    id = request.forms.app_id
    title = request.forms.title
    description = request.forms.description
    min_number = request.forms.min_number
    max_number = request.forms.max_number
    memory = request.forms.memory

    if user_id == 0:
        return abort(403, '未登录！')

    dao = AppModel.update(title=title, description=description, min_container_number=min_number,
                          max_container_number=max_number, memory=memory).where(AppModel.id == id,
                                                                                AppModel.user_id == user_id)
    dao.execute()

    return redirect("/user")


@user_app.get("/account")
def account():
    """
    账号页面
    :return:
    """
    return __output_html("account")


@user_app.post("/login")
def login():
    """
    用户登录
    :return:
    """

    username = request.forms.username
    password = request.forms.password
    password = hashlib.md5(password.encode("UTF-8")).hexdigest()

    try:
        user = UserModel.select().where(UserModel.username == username, UserModel.password == password, UserModel.status == 0).get()
        request.session['uid'] = user.id
    except:
        return {"code": -1}

    return {"code": 0}


@user_app.post("/deploymentApp")
def deployment_app():
    """
    重新部署应用
    :return:
    """

    user_id = get_login_user_id()
    app_id = request.forms.app_id
    # 防止越权操作
    AppModel.select().where(AppModel.user_id == user_id).get()

    if TaskQueueModel.select().where(TaskQueueModel.app_id == app_id, TaskQueueModel.command_code == 3).count():
        return {"code": -1}

    obj = TaskQueueModel(app_id=app_id, user_id=user_id, command_code=3, command_content='{}', create_time=datetime.datetime.now())
    obj.save()

    return {"code": 0}