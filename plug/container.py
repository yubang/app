# coding:UTF-8


"""
docker容器管理插件
@author: yubang
20151124
"""


from bottle import Bottle, request
from model.container import ContainerModel
from lib import docker
from lib.config import get_config_data
import subprocess
import datetime
import urllib


container_app = Bottle()


def check_token(token):
    """
    检测token凭证
    :param token: 登录token
    :return: boolean
    """
    data = get_config_data()
    return token == data['token.token']


def get_one_able_post():
    """
    获取一个可用端口
    :return: int
    """
    port = 10000
    objs = ContainerModel.select()
    while port < 50000:
        sign = True
        for obj in objs:
            if obj.port == port:
                sign = False
                break
        if sign:
            return port
        port += 1
    return port


def get_realname_from_image_name(image_name):
    """
    获取系统镜像名字（转义处理镜像名字）
    :param image_name:
    :return: str
    """
    system_images = dict()
    system_images['static'] = 'paas-static'
    system_images['python2.7'] = 'paas-python2'
    system_images['python3.5'] = 'paas-python3'
    system_images['go1.5'] = 'paas-go'
    system_images['java8'] = 'paas-java8'
    system_images['nodejs4'] = 'paas-nodejs4'
    return system_images.get(image_name, 'paas-static')


def build_one_container(image_name, memory, code_address):
    """
    生成一个容器
    :param image_name: 镜像名字
    :param memory: 内存
    :param code_address: 代码zip地址
    :return: dict
    """

    port = get_one_able_post()
    system_image = get_realname_from_image_name(image_name)
    command = "docker run -d -m %dm -p %d:80 %s /bin/bash /tmp/start.sh '%s'" % (int(memory), port, system_image, code_address)

    subprocess.getstatusoutput(command)

    code, result = subprocess.getstatusoutput("docker ps | grep -v grep |grep 0.0.0.0:%d|awk -F ' ' '{print $1}'" % port)
    if code != 0:
        return {"code": 10004}

    # 记录新建容器端口情况
    obj = ContainerModel(port=port, container_id=result, memory=memory, code_address=code_address,
                         image_name=system_image, create_time=datetime.datetime.now())
    obj.save()

    return {"code": 0, "result": {"containerId": result, "port": port}}


def option_container(container_id, token, option):
    """
    操作容器
    :param container_id: 容器id
    :param token: 校验token
    :param option: 操作命令
    :return:
    """
    if not check_token(token):
        return {"code": 10001}

    if not container_id:
        return {"code": 10002}

    command = "docker %s %s" % (option, container_id)
    code, _ = subprocess.getstatusoutput(command)
    if code != 0:
        return {"code": 10004}

    return {"code": 0}


@container_app.post("/build")
def build():
    """
    生成一个容器
    :return:
    """

    image_name = request.forms.get('image_name', None)
    code_address = request.forms.get('code_address', None)
    memory = request.forms.get('memory', None)
    token = request.forms.get('token', None)

    # 特殊处理
    code_address = urllib.parse.unquote(code_address)

    if not check_token(token):
        return {"code": 10001}

    if not image_name or not code_address or not memory or not token:
        return {"code": 10002}

    return build_one_container(image_name, memory, code_address)


@container_app.post("/start")
def start():
    """
    启动容器
    :return:
    """

    container_id = request.forms.get('containerId', None)
    token = request.forms.get('token', None)

    return option_container(container_id, token, "start")


@container_app.post("/restart")
def restart():
    """
    启动容器
    :return:
    """

    container_id = request.forms.get('containerId', None)
    token = request.forms.get('token', None)

    return option_container(container_id, token, "restart")


@container_app.post("/stop")
def stop():
    """
    停止容器
    :return:
    """

    container_id = request.forms.get('containerId', None)
    token = request.forms.get('token', None)

    return option_container(container_id, token, "stop")


@container_app.post("/remove")
def remove():
    """
    删除容器
    :return:
    """

    container_id = request.forms.get('containerId', None)
    token = request.forms.get('token', None)

    r = option_container(container_id, token, "rm -f")

    if r['code'] == 0:
        dao = ContainerModel.delete().filter(container_id=container_id)
        dao.execute()
    return r


@container_app.post("/analy")
def analy():
    """
    获取主机容器状态
    :return: dict
    """

    token = request.forms.get('token', None)
    if not check_token(token):
        return {"code": 10001}

    count = ContainerModel.select().count()
    memory = ContainerModel.get_total_memory()

    return {"code": 0, "result": {"count": count, "totalMemory": memory}}


@container_app.post("/stats")
def stats():
    """
    获取容器cpu和内存
    :return: dict
    """

    container_id =request.forms.get('containerId', None)
    token = request.forms.get('token', None)

    if not check_token(token):
        return {"code": 10001}

    if not container_id:
        return {"code": 10002}

    code, r = docker.get_container_memory_and_cpu(container_id)

    if code != 0:
        return {"code": 10004}

    return {"code": 0, "result": {"memory": r['memory'], "cpu": r['cpu']}}
