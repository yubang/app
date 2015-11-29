# coding:UTF-8


"""
后台管理界面模块
@author: yubang
"""


from bottle import Bottle, static_file, request, response, redirect
from model.user import UserModel
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
