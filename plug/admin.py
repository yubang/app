# coding:UTF-8


"""
后台管理界面模块
@author: yubang
"""


from bottle import Bottle, static_file, request, response, redirect
from model.user import UserModel
from model.container_server import ContainerServerModel
import os
import datetime
import hashlib
import json

admin_app = Bottle()


def __output_html(html_name):
    """
    输出html文件
    :param html_name: 文件名（不要后缀）
    :return:
    """
    path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.dirname(path)
    path = os.path.join(path, "static/html/admin")
    return static_file(html_name+".html", root=path)


@admin_app.get("/")
def index():
    """
    后台主页面
    :return:
    """
    return __output_html("index")


@admin_app.get("/user")
def user():
    """
    用户管理
    :return:
    """
    return __output_html("user")


@admin_app.post("/addUser")
def add_user():
    """
    添加新用户
    :return:
    """

    username = request.forms.username
    password = request.forms.password
    nickname = request.forms.nickname
    status = request.forms.status
    create_time = datetime.datetime.now()
    # 加密密码
    password = hashlib.md5(password.encode("UTF-8")).hexdigest()

    # 检测用户名是否存在
    if UserModel.select().where(UserModel.username == username, UserModel.status != 2).count():
        return {"code": -1}

    user = UserModel(username=username, password=password, nickname=nickname, status=status, create_time=create_time)
    user.save()

    return {"code": 0}


@admin_app.get("/getUsers")
def get_users():
    """
    获取所有用户
    :return:
    """
    response.set_header("Content-Type", "application/json")
    users = UserModel.select().where(UserModel.status != 2).order_by(UserModel.id.desc())
    r = list(map(UserModel.get_dict_from_obj, users))
    return json.dumps(r)


@admin_app.get("/deleteUser/:user_id#\d+#")
def delete_user(user_id):
    """
    删除一个用户
    :param user_id: 用户id
    :return:
    """

    dao = UserModel.update(status=2).where(UserModel.id == user_id)
    dao.execute()

    return redirect("/admin/user")


@admin_app.post("/updateUser")
def update_user():
    """
    修改用户信息
    :return:
    """
    user_id = request.forms.user_id
    username = request.forms.username
    password = request.forms.password
    nickname = request.forms.nickname
    status = request.forms.status

    # 判断是否需要修改密码
    if password == '':
        query = UserModel.update(username=username, nickname=nickname, status=status).where(UserModel.id == user_id)
    else:
        # 加密密码
        password = hashlib.md5(password.encode("UTF-8")).hexdigest()
        query = UserModel.update(username=username, nickname=nickname, status=status, password=password).where(UserModel.id == user_id)

    if query.execute():
        return {"code": 0}

    return {"code": -1}


@admin_app.get("/server")
def server():
    """
    获取容器服务器配置
    :return:
    """
    return __output_html("server")


@admin_app.get("/getServers")
def get_servers():
    """
    获取所有的容器服务器
    :return:
    """
    lists = ContainerServerModel.select().where(ContainerServerModel.status != 2).order_by(ContainerServerModel.sort.desc(), ContainerServerModel.id.desc())
    objs = list(map(ContainerServerModel.get_dict_from_obj, lists))
    
    return json.dumps(objs)


@admin_app.post("/addServer")
def add_server():
    """
    添加容器服务器
    :return:
    """
    title = request.forms.title
    server_host = request.forms.server_host
    server_port = request.forms.server_port
    status = request.forms.status
    max_container_number = request.forms.max_container_number
    max_memory = request.forms.max_memory
    sort = request.forms.sort
    create_time = datetime.datetime.now()

    dao = ContainerServerModel(server_host=server_host, server_port=server_port, status=status, title=title,
                               max_container_number=max_container_number, max_memory=max_memory, sort=sort, create_time=create_time)

    if not dao.save():
        return {"code": -1}

    return {"code": 0}