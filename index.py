# coding:UTF-8


from bottle import run, Bottle
from plug.container import container_app
from plug.user import user_app
from plug.admin import admin_app
from plug.code import code_app
from model.base import db
from lib.middle import MiddleSupport
from middle.static import StaticMiddle
from middle.session import SessionMiddle
from middle.time_middle import TimeMiddle
from middle.power import PowerMiddle
from model.base import start_connect, close_connect

app = Bottle()

app.mount("/container", container_app)
app.mount("/user", user_app)
app.mount("/admin", admin_app)
app.mount("/code", code_app)


@app.hook("before_request")
def __db_connect():
    start_connect()


@app.hook("after_request")
def __db_close():
    close_connect()


app = MiddleSupport(app)
app.add_middle_plug(SessionMiddle)
app.add_middle_plug(StaticMiddle)
app.add_middle_plug(TimeMiddle)
app.add_middle_plug(PowerMiddle)

run(app=app, host="0.0.0.0", port=8000, reloader=True, debug=True)

