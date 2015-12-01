# coding:UTF-8


"""
封装自定义session中间件
@author: yubang
"""


from bottle import request


class SessionMiddle(object):
    def before_request(self, request, response):
        request.session = {}
        return None
    def destroy(self):
        request.session = {}
        return None