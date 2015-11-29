# coding:UTF-8


"""
用户模型
@author: yubang
"""


from peewee import CharField, IntegerField, DateTimeField
from model.base import BaseModel


class UserModel(BaseModel):
    """
    用户模型类
    """

    class Meta:
        db_table = "user"

    id = IntegerField()
    username = CharField(max_length=20)
    password = CharField(max_length=32)
    nickname = CharField(max_length=20)
    status = IntegerField()
    create_time = DateTimeField()

    @classmethod
    def get_dict_from_obj(cls, user_obj):
        r = dict()
        r['id'] = user_obj.id
        r['username'] = user_obj.username
        r['status'] = user_obj.status
        r['nickname'] = user_obj.nickname
        r['create_time'] = user_obj.create_time.strftime("%Y-%m-%d %H:%M:%S")
        return r

