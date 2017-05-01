package web

import "../ctsFrame/cacheTools"

type OwnConfigInfo struct {
	HttpAddr string
	RedisObject cacheTools.RedisClientObject
}
