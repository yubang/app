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

/*
容器服务器相关
 */
// 一个记录容器服务器报告在线的hash
var REDIS_KEY_CONTAINER_SERVER_LAST_TIME_HASH = "paas_container_server_last_time_hash"
// 一个确定唯一性的容器服务器set
var REDIS_KEY_CONTAINER_SERVER_IP_SET = "paas_container_server_ip_set"
// 一个排序的容器服务器zset
var REDIS_KEY_CONTAINER_SERVER_IP_ZSET = "paas_container_server_ip_zset"
// 一个排序的容器服务器list
var REDIS_KEY_CONTAINER_SERVER_IP_LIST = "paas_container_server_ip_list"