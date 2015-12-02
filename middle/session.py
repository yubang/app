# coding:UTF-8


"""
封装自定义session中间件
@author: yubang
"""

import hashlib
import datetime
import os
import pickle
import time


def get_session_file(session_id):
    dir_name = os.path.dirname(os.path.dirname(os.path.realpath(__file__)))
    dir_name = os.path.join(dir_name, "data", "session")
    session_id = hashlib.md5(session_id.encode("UTF-8")).hexdigest()
    return os.path.join(dir_name, session_id)


def get_session(session_id):
    fp_path = get_session_file(session_id)
    if not os.path.exists(fp_path):
        return {}
    fp = open(fp_path, "rb")
    data = fp.read()
    fp.close()

    obj = pickle.loads(data)
    if obj['time'] < time.time():
        os.remove(fp_path)
        return {}
    return obj['data']


def set_session(session_id, session_value, timeout):
    fp = open(get_session_file(session_id), "wb")
    obj = {"data": session_value, "time": time.time() + timeout}
    obj = pickle.dumps(obj)
    fp.write(obj)
    fp.close()


class SessionMiddle(object):
    def before_request(self, request, response):
        self.request = request

        # 先判断有没有session_id的cookie
        session_id = request.cookies.get('session_id', None)
        if session_id is None:
            session_id = hashlib.md5(datetime.datetime.now().strftime("%Y%m%d%H%M%S").encode("UTF-8")).hexdigest()
            response.set_cookie("session_id", session_id)
        self.__session_id = session_id
        request.session = get_session(session_id)
        return None
    def destroy(self):
        set_session(self.__session_id, self.request.session, 3600)
        self.request.session = {}
        return None