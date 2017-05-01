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
