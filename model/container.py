# coding:UTF-8


"""
单机容器信息模型
@author: yubang
"""

from peewee import CharField, IntegerField, DateTimeField
from model.base import BaseModel


class ContainerModel(BaseModel):
    """
    容器表模型
    """
    class Meta:
        db_table = "container"

    id = IntegerField()
    container_id = CharField(max_length=100)
    port = IntegerField()
    code_address = CharField(max_length=255)
    memory = IntegerField()
    image_name = CharField(max_length=20)
    create_time = DateTimeField()