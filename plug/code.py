# coding:UTF-8


"""
简易代码管理模块
@author: yubang
"""

from bottle import Bottle, static_file, response, request, redirect, abort
from model.code import CodeModel
import os
import json
import datetime
import hashlib


code_app = Bottle()


def __output_html(html_name):
    """
    输出html文件
    :param html_name: 文件名（不要后缀）
    :return:
    """
    path = os.path.dirname(os.path.realpath(__file__))
    path = os.path.dirname(path)
    path = os.path.join(path, "static/html/code")
    return static_file(html_name+".html", root=path)


def get_login_user_id():
    """
    获取登录用户id，未登录返回0
    :return: int
    """
    return 1


@code_app.get("/")
def index():
    return __output_html("index")


@code_app.get("/getAllCode")
def get_all_codes():
    """
    获取所有代码仓库
    :return:
    """
    user_id = get_login_user_id()
    lists = CodeModel.select().where(CodeModel.status != 2, CodeModel.user_id == user_id)
    objs = list(map(CodeModel.get_dict_from_obj, lists))

    response.set_header("Content-Type", "application/json")
    return json.dumps(objs)


@code_app.post("/addCode")
def add_code():
    """
    添加代码仓库
    :return:
    """

    title = request.forms.title
    status = request.forms.status
    user_id = get_login_user_id()

    create_time = datetime.datetime.now()
    token = hashlib.md5(create_time.strftime("%Y%m%d%H%M%S").encode("UTF-8")).hexdigest()
    warehouse = hashlib.md5(str(user_id).encode("UTF-8")).hexdigest()

    dao = CodeModel(title=title, status=status, user_id=user_id, create_time=create_time, token=token,
                    warehouse=warehouse)
    if not dao.save():
        return {"code": -1}

    return {"code": 0}


@code_app.get("/deleteCode/:code_id#\d+#")
def delete_code(code_id):
    """
    删除代码仓库
    :param code_id: 仓库id
    :return:
    """

    user_id = get_login_user_id()

    dao = CodeModel.update(status=2).where(CodeModel.user_id == user_id, CodeModel.id == code_id)
    dao.execute()

    return redirect("/code")


@code_app.post("/updateCode")
def update_code():
    """
    跟新代码仓库
    :return:
    """

    id = request.forms.id
    user_id = get_login_user_id()
    title = request.forms.title
    status = request.forms.status

    dao = CodeModel.update(title=title, status=status).where(CodeModel.id == id, CodeModel.user_id == user_id)
    if not dao.execute():
        return {"code": -1}

    return {"code": 0}


@code_app.get("/download/:warehouse#[0-9a-z]{32}#/:token#[0-9a-z]{32}#")
def download(warehouse, token):
    """
    下载代码
    :param warehouse: 仓库名称
    :param token: 下载凭证
    :return:
    """

    try:
        obj = CodeModel.select().where(CodeModel.warehouse == warehouse, CodeModel.token == token, CodeModel.status == 0).get()
    except Exception:
        return abort(404)

    fp_path = os.path.dirname(os.path.dirname(os.path.realpath(__file__))) + "/data/code"
    return static_file(str(obj.id)+".zip", root=fp_path, download=True)


@code_app.post("/uploadFile")
def upload_zip():
    """
    上传文件
    :return:
    """

    code_id = int(request.forms.code_id)
    obj = CodeModel.select().where(CodeModel.user_id == get_login_user_id(), CodeModel.id == code_id).get()

    fp = request.files['file']
    fp_path = os.path.dirname(os.path.dirname(os.path.realpath(__file__))) + "/data/code/"
    fp.save(fp_path + str(obj.id) + ".zip", overwrite=True)

    return "ok"
