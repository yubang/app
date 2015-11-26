# coding:UTF-8


"""
封装docker操作
@author: yubang
"""


import socket
import json


def __socket_to_docker(docker_command):
    """
    与docker交互
    :param docker_command: docker API 指令
    :return: 结果字典
    """
    sock = socket.socket(socket.AF_UNIX, socket.SOCK_STREAM)
    sock.settimeout(5)
    sock.connect("/var/run/docker.sock")
    sock.send(('GET %s HTTP/1.1\r\n' % docker_command).encode())
    sock.send('\r\n'.encode())

    r = sock.recv(10000)
    r = r.decode("UTF-8")
    r = r.split("\r\n")
    r = r[len(r)-2]

    result_obj = json.loads(r)
    sock.close()
    return result_obj


def docker_get_message(docker_command):
    try:
        return 0, __socket_to_docker(docker_command)
    except Exception:
        raise
        return -1, None

def get_container_memory_and_cpu(container_id):
    """
    获取容器内存与cpu使用率
    :param container_id: 容器id
    :return: 元组（cpu,内存）
    """
    docker_command = "/containers/%s/stats?stream=0" % container_id
    code, obj = docker_get_message(docker_command)

    if code != 0:
        return code, {"cpu": 0, "memory": 0}

    memory = int(obj['memory_stats']['usage']) * 100 / int(obj['memory_stats']['limit'])
    memory = float('%0.2f' % memory)

    return 0, {"cpu": 0, "memory": memory}
