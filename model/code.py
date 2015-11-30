# coding:UTF-8


"""
代码仓库模型
@author: yubang
"""


from model.base import BaseModel
from peewee import CharField, DateTimeField, IntegerField


class CodeModel(BaseModel):
    class Meta:
        db_table = "code"
    id = IntegerField()
    user_id = IntegerField()
    title = CharField()
    token = CharField(max_length=32)
    warehouse = CharField(max_length=32)
    status = IntegerField()
    create_time = DateTimeField()

    @classmethod
    def get_dict_from_obj(cls, obj):
        r = dict()
        r['id'] = obj.id
        r['title'] = obj.title
        r['token'] = obj.token
        r['warehouse'] = obj.warehouse
        r['status'] = obj.status
        r['create_time'] = obj.create_time.strftime("%Y-%m-%d %H:%M:%S")
        return r
