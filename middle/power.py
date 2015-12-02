# coding:UTF-8


"""
权限控制中间件
@author: yubang
"""


import re


class PowerMiddle(object):
    def before_request(self, request, response):
        url = request.path
        if re.search(r'^/user', url) and url != '/user/account' and 'uid' not in request.session:
            response.redirect("/user/account")
            return response