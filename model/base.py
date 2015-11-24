# coding:UTF-8


"""
模型基类
@author: yubang
"""

from peewee import Model, SqliteDatabase


db = SqliteDatabase("db/base.db")


class BaseModel(Model):
    class Meta:
        database = db
