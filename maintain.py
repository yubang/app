# coding:UTF-8


"""
自动维护脚本
@author: yubang
"""


from model.base import start_connect, close_connect, db
from model.app_container import AppContainModel
from script.task_execute import request_api
from lib.log import login_log
import os
import json
import pickle
import time


path = os.path.dirname(os.path.realpath(__file__))


# 第一步清除所有过期的session文件
print("开始清除session")
session_path = os.path.join(path, "data", "session")
fps = os.listdir(session_path)
for fp in fps:
    fp_path = session_path + "/" + fp
    fp = open(fp_path, "rb")
    data = fp.read()
    fp.close()
    obj = pickle.loads(data)
    if obj['time'] < time.time():
        os.remove(fp_path)


print("清除session结束")


# 清除孤儿容器
print("开始清除孤儿容器")
start_connect(sqlite_db_use=False)
c = db.execute_sql("select id, api_url, container_id, app_id, host, port  from app_container where app_id not in "
                   "(select id from app)")
objs = c.fetchall()
for t in objs:
    obj['id'] = t[0]
    obj['api_url'] = t[1]
    obj['container_id'] = t[2]
    obj['app_id'] = t[3]
    obj['host'] = t[4]
    obj['port'] = t[5]
    login_log("maintain", "发现孤儿容器，所属应用id：%d，容器id：%s，容器域名：%s，容器端口：%d" % (obj['app_id'],
                                                                        obj['container_id'], obj['host'], obj['port']))
    api_url = obj['api_url'] + "remove"
    data = {
        "containerId": obj['container_id']
    }
    r = request_api(api_url, data)
    if r.status_code == 200:
        d = json.loads(r.text)
        if d['code'] == 0:
            # 删除容器成功
            dao = AppContainModel.delete().where(AppContainModel.id == obj['id'])
            dao.execute()
            login_log("maintain", "删除孤儿容器成功，所属应用id：%d，容器id：%s，容器域名：%s，容器端口：%d" % (obj['app_id'],
                                                                                  obj['container_id'], obj['host'],
                                                                                  obj['port']))
            continue
    login_log("maintain_error", "删除孤儿容器失败，所属应用id：%d，容器id：%s，容器域名：%s，容器端口：%d" % (obj['app_id'],
                                                                                obj['container_id'], obj['host'],
                                                                                obj['port']))

close_connect(sqlite_db_use=False)
print("清除孤儿容器结束")
