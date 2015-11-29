# coding:UTF-8


from bottle import run, static_file, Bottle
from plug.container import container_app
from plug.user import user_app
from plug.admin import admin_app
from model.base import db
import os


app = Bottle()
app.mount("/container", container_app)
app.mount("/user", user_app)
app.mount("/admin", admin_app)


@app.route("/")
def index():
    return static_file("index.html", root=os.path.dirname(os.path.realpath(__file__))+"/static/html")


@app.route("/static/css/<filename>")
def css(filename):
    return static_file(filename, root=os.path.dirname(os.path.realpath(__file__))+"/static/css")


@app.route("/static/js/<filename>")
def js(filename):
    return static_file(filename, root=os.path.dirname(os.path.realpath(__file__))+"/static/js")


@app.hook("before_request")
def __db_connect():
    db.connect()


@app.hook("after_request")
def __db_close():
    if not db.is_closed():
        db.close()


run(app=app, host="0.0.0.0", port=8000, reloader=True, debug=True)

