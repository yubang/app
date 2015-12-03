# coding:UTF-8

"""
应用模型
@author: yubang
"""

from peewee import CharField, IntegerField, DateTimeField
from model.base import BaseModel
from model.app_container import AppContainModel


class AppModel(BaseModel):
    class Meta:
        db_table = "app"
    id = IntegerField()
    user_id = IntegerField()
    title = CharField(max_length=20)
    description = CharField(max_length=200)
    memory = IntegerField()
    env = CharField(max_length=20)
    code_address = CharField(max_length=255)
    app_host = CharField(max_length=200)
    app_port = IntegerField()
    min_container_number = IntegerField()
    max_container_number = IntegerField()
    create_time = DateTimeField()

    @classmethod
    def get_dict_from_model_obj(cls, model_obj):
        r = dict()
        r['id'] = model_obj.id
        r['title'] = model_obj.title
        r['description'] = model_obj.description
        r['memory'] = model_obj.memory
        r['env'] = model_obj.env
        r['code_address'] = model_obj.code_address
        r['app_host'] = model_obj.app_host
        r['app_port'] = model_obj.app_port
        r['min_container_number'] = model_obj.min_container_number
        r['max_container_number'] = model_obj.max_container_number
        r['create_time'] = model_obj.create_time.strftime("%Y-%m-%d %H:%M:%S")
        r['now_container_number'] = AppContainModel.select().where(AppContainModel.app_id == model_obj.id).count()
        return r
