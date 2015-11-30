# coding:UTF-8


"""
容器服务器模型
"""


from peewee import CharField, IntegerField, DateTimeField
from model.base import BaseModel


class ContainerServerModel(BaseModel):
    class Meta:
        db_table = "container_server"
    id = IntegerField()
    title = CharField(max_length=20)
    server_host = CharField(max_length=200)
    server_port = IntegerField()
    status = IntegerField()
    max_container_number = IntegerField()
    max_memory = IntegerField()
    sort = IntegerField()
    create_time = DateTimeField()

    @classmethod
    def get_dict_from_obj(cls, obj):
        r = dict()
        r['id'] = obj.id
        r['title'] = obj.title
        r['server_host'] = obj.server_host
        r['server_port'] = obj.server_port
        r['status'] = obj.status
        r['max_container_number'] = obj.max_container_number
        r['max_memory'] = obj.max_memory
        r['sort'] = obj.sort
        r['create_time'] = obj.create_time.strftime("%Y-%m-%d %H:%M:%D")
        return r
