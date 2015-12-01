# coding:UTF-8


"""
封装自定义session中间件
@author: yubang
"""

import hashlib
import datetime


class SessionMiddle(object):
    def before_request(self, request, response):
        self.request = request
        request.session = {}

        # 先判断有没有session_id的cookie
        session_id = request.cookies.get('session_id', None)
        if session_id is None:
            session_id = hashlib.md5(datetime.datetime.now().strftime("%Y%m%d%H%M%S").encode("UTF-8")).hexdigest()
            response.set_cookie("session_id", session_id)

        return None
    def destroy(self):
        self.request.session = {}
        return None