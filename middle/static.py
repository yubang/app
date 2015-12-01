# coding:UTF-8


"""
处理静态资源文件中间件
@author: yubang
"""


from bottle import static_file, redirect
import re
import os


class StaticMiddle(object):
    def before_request(self, request, response):
        url = request.fullpath
        if re.search(r'^/static/', url):
            dir_path = os.path.dirname(os.path.dirname(os.path.realpath(__file__)))
            dir_path = dir_path + url
            file_name = os.path.basename(dir_path)
            dir_path = os.path.dirname(dir_path)
            obj = static_file(file_name, root=dir_path)
            if type(obj.body).__name__ == 'str':
                response.body = obj.body.encode()
            else:
                response.body = obj.body.read()

            response.status = obj.status
            for obj in obj.headerlist:
                response.add_header(obj[0], obj[1])
            return response
        return None
    def destroy(self):
        return None