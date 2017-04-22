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
// 记录容器服务器具体信息hash（请自行添加ip）
var REDIS_KEY_CONTAINER_SERVER_INFO_HASH = "paas_container_server_info_hash_"

/*
应用相关
 */

// APP信息hash（请自行添加appid）
var REDIS_KEY_APP_MESSAGE_HASH = "paas_app_message_hash_"

// APP列表
var REDIS_KEY_APP_LIST = "paas_app_list"

/*
容器服务器任务相关
 */
// 一个计划应用分配json字符串（请自行添加）
var REDIS_KEY_PLAN_CONTAIN_USE_STR = "paas_plan_container_use_"

/*
反向代理相关
 */
// 记录应用与后端ip端口的set（请自行加上appid）
var REDIS_KEY_PROXY_APP_CONTAINER_IP_AND_PORT_SET = "paas_proxy_app_container_ip_and_port_set"
// 记录域名与app绑定关系
var REDIS_KEY_PROXY_DOMAIN_APP_HASH = "paas_proxy_domain_app_hash"
// 记录当前容器服务器应用与容器映射关系
var REDIS_KEY_PROXY_APP_CONTAINER_HASH = "paas_proxy_app_container_hash"
