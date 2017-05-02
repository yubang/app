package web

import "../ctsFrame/cacheTools"

type OwnConfigInfo struct {
	HttpAddr string
	RedisObject cacheTools.RedisClientObject
}

// REDIS KEY
var REDIS_KEY_APP_ZSET = "paas_app_zset"
var REDIS_KEY_APP_INFO_HSET = "paas_app_info_hset_"
var REDIS_KEY_APP_LOG_LIST = "paas_app_log_list_"

var REDIS_KEY_APP_DOMAIN_HSET = "paas_app_domain"
var REDIS_KEY_DOMAIN_APP_HSET = "paas_domain_app"

var REDIS_KEY_POST_USE = "paas_port_use_set"

var REDIS_KEY_APP_IMAGE_LIST = "paas_app_image_list_"