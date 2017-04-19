package config

import (
	"../../paas/tools"
	"net/http"
)

type RedisConfig struct {
	Host string
	Port int
	Db int
	Password string
}

type HttpProxyConfig struct {
	Ip string
	Port float64
}

type ApiServerConfig struct {
	Ip string
	Port float64
}

type PassConfig struct {
	Err bool
	Token string
	HttpProxyConfigData HttpProxyConfig
	ApiServerConfigData ApiServerConfig
	RedisConfigData RedisConfig
}

func GetPaasConfig() PassConfig{
	jsonString := tools.ReadFromFile("./config.json")
	obj := tools.JsonToInterface(jsonString)

	passConfig := new(PassConfig)

	if obj == nil{
		passConfig.Err = true
		return *passConfig
	}else {
		passConfig.Err = false
	}

	passConfig.Token = obj["token"].(string)

	httpProxyConfig := new(HttpProxyConfig)
	httpProxyConfig.Ip = obj["httpProxy"].(map[string]interface{})["ip"].(string)
	httpProxyConfig.Port = obj["httpProxy"].(map[string]interface{})["port"].(float64)
	passConfig.HttpProxyConfigData = *httpProxyConfig

	apiServerConfig := new(ApiServerConfig)
	apiServerConfig.Ip = obj["apiServer"].(map[string]interface{})["ip"].(string)
	apiServerConfig.Port = obj["apiServer"].(map[string]interface{})["port"].(float64)
	passConfig.ApiServerConfigData = *apiServerConfig

	redisConfig := new(RedisConfig)
	redisConfig.Host = obj["redis"].(map[string]interface{})["host"].(string)
	redisConfig.Port = tools.Float64ToInt(obj["redis"].(map[string]interface{})["port"].(float64))
	redisConfig.Db = tools.Float64ToInt(obj["redis"].(map[string]interface{})["db"].(float64))
	redisConfig.Password = obj["redis"].(map[string]interface{})["password"].(string)
	passConfig.RedisConfigData = *redisConfig

	return *passConfig
}

// 获取paas平台名字
func GetPaasServerName() string{
	return "paas 0.1.0"
}

// 输出paas平台名字
func OutputPaasName(w http.ResponseWriter){
	w.Header().Set("paas-server", GetPaasServerName())
}