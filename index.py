# coding:UTF-8


from bottle import route, run, static_file
import os


@route("/")
def index():
    return static_file("index.html", root=os.path.dirname(os.path.realpath(__file__))+"/static/html")


@route("/static/css/<filename>")
def css(filename):
    return static_file(filename, root=os.path.dirname(os.path.realpath(__file__))+"/static/css")


@route("/static/js/<filename>")
def js(filename):
    return static_file(filename, root=os.path.dirname(os.path.realpath(__file__))+"/static/js")



run(host="0.0.0.0", port=8000, reloader=True, debug=True)
