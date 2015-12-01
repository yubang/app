# coding:UTF-8


"""
bottle中间件支持
@author: yubang
"""


from bottle import Request, Response, response as bottle_response


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
        # 请求前处理
        request = Request(environ)
        response = MiddleResponse()
        r = self.__use_plug_before_request(request, response)
        if r:
            return self.__output(environ, start_response, r)

        # 调用bottle处理
        out = self.__app._cast(self.__app._handle(environ))
        # 提取bottle处理后的结果
        response.status = bottle_response.status
        for obj in bottle_response.headerlist:
            response.add_header(obj[0], obj[1])

        for line in out:
            response.body += line

        if hasattr(out, 'close'):
            out.close()

        # 请求后处理
        response = self.__use_plug_after_request(request, response)

        # 销毁插件
        self.__handle_plug_in_end()

        # 输出结果
        return self.__output(environ, start_response, response)


    def __output(self, environ, start_response, response):
        response.set_header("Content-Length", len(response.body))
        start_response(response.status, response.headerlist)
        return [response.body]

    def __use_plug_before_request(self, request, response):
        """
        在访问前使用插件处理
        :return:
        """

        for plug in self.__middle_plug:
            if not hasattr(plug, "before_request"):
                continue
            r = plug.before_request(request, response)
            if r:
                return r
        return None

    def __use_plug_after_request(self, request, response):
        self.__middle_plug.reverse()

        for plug in self.__middle_plug:
            if hasattr(plug, "after_request"):
                response = plug.after_request(request, response)

        self.__middle_plug.reverse()
        return response

    def __handle_plug_in_end(self):
        for plug in self.__middle_plug:
            if hasattr(plug, "destroy"):
                r = plug.destroy()

    def add_middle_plug(self, middle_plug):
        self.__middle_plug.append(middle_plug())
