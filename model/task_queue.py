# coding:UTF-8


"""
任务队列模型
@author: yubang
"""


from model.base import BaseModel
from peewee import IntegerField, DateTimeField, TextField


class TaskQueueModel(BaseModel):
    class Meta:
        db_table = "task_queue"
    id = IntegerField()
    user_id = IntegerField()
    app_id = IntegerField()
    command_code = IntegerField()
    command_content = TextField()
    callback_sql = TextField()
    callback_sql_error = TextField()
    create_time = DateTimeField()
