# coding:UTF-8


"""
bottle中间件支持
@author: yubang
"""


from bottle import Request, Response


class MiddleResponse(Response):
    """
    中间件response类
    """
    def __init__(self):
        super().__init__()
        self.status = '200 ok'
        self.body = b''
    def redirect(self, url):
        self.body = b''
        self.set_header('Location', url)
        self.status = '302 location'


class MiddleSupport(object):
    """
    中间件支持类
    """

    def __init__(self, app):
        self.__app = app
        self.__middle_plug = []

    def __call__(self, environ, start_response):
        request = Request(environ)
        response = MiddleResponse()
        r = self.__use_plug_before_request(request, response)
        if r:
            return self.__output(environ, start_response, r)

        r = self.__app(environ, start_response)

        self.__handle_plug_in_end()

        return r

    def __output(self, environ, start_response, response):
        start_response(response.status, response.headerlist)
        return [response.body]

    def __use_plug_before_request(self, request, response):
        """
        在访问前使用插件处理
        :return:
        """

        for plug in self.__middle_plug:
            r = plug.before_request(request, response)
            if r:
                return r
        return None

    def __handle_plug_in_end(self):
        for plug in self.__middle_plug:
            r = plug.destroy()

    def add_middle_plug(self, middle_plug):
        self.__middle_plug.append(middle_plug())
