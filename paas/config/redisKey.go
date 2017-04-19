package config

/*
redis相关key配置
@author: yubang
 */

// 任务池（请自行添加当前日期）
var REDIS_KEY_TASK_HASH = "paas_task_hash_"

// 构建镜像任务队列
var REDIS_KEY_TASK_IMAGE_LIST = "paas_task_image_list"

// 任务处理后续任务队列
var REDIS_KEY_TASK_AFTER_HANDLER_LIST = "paas_task_after_handler_list"

// APP信息hash（请自行添加appid）
var REDIS_KEY_APP_MESSAGE_HASH = "paas_app_message_hash_"