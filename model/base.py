# coding:UTF-8


"""
模型基类
@author: yubang
"""

from peewee import Model, SqliteDatabase, MySQLDatabase
from lib.config import get_config_data

d = get_config_data()
db = MySQLDatabase(database=d['mysql.db_name'], user=d['mysql.db_username'], passwd=d['mysql.db_password'],
                   host=d['mysql.db_host'], port=int(d['mysql.db_port']), charset="utf8")
sqlite_db = SqliteDatabase("data/db/base.db", threadlocals=True)


def start_connect():
    db.connect()
    sqlite_db.connect()


def close_connect():
    if not db.is_closed():
        db.close()
    if not sqlite_db.is_closed():
        sqlite_db.close()


class BaseModel(Model):
    class Meta:
        database = db