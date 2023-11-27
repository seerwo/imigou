package config

import "github.com/seerwo/imigou/cache"

//Config config for wms
type Config struct {
	AppID string `json:"app_id"`				//appid
	AppSecret string `json:"app_secret"`		//appsecret
	Token string `json:"token"`					//token
	EncodingAESKey string `json:"encoding_aes_key"`//EncodingAESKey
	Cache cache.Cache
}