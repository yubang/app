# coding:UTF-8


"""
应用容器记录模型
@author: yubang
"""


from model.base import BaseModel
from peewee import CharField, IntegerField, DateTimeField


class AppContainModel(BaseModel):
    class Meta:
        db_table = "app_container"
    id = IntegerField()
    app_id = IntegerField()
    host = CharField(max_length=100)
    port = IntegerField()
    create_time = DateTimeField()
