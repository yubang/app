# coding:UTF-8


"""
单机容器信息模型
@author: yubang
"""

from peewee import CharField, IntegerField, DateTimeField
from model.base import BaseModel, db, sqlite_db


class ContainerModel(BaseModel):
    """
    容器表模型
    """
    class Meta:
        database = sqlite_db
        db_table = "container"

    @classmethod
    def get_total_memory(cls):
        """
        获取容器占用内存
        :return: int
        """
        c = db.execute_sql("select sum(memory) from container")
        obj = c.fetchone()
        if not obj or not obj[0]:
            return 0
        return int(obj[0])

    id = IntegerField()
    container_id = CharField(max_length=100)
    port = IntegerField()
    code_address = CharField(max_length=255)
    memory = IntegerField()
    image_name = CharField(max_length=20)
    create_time = DateTimeField()
