package config

import "github.com/amorist/douyin/open/cache"

//Config config for 抖音开放平台
type Config struct {
	ClientKey    string `json:"client_key"`    // client_key
	ClientSecret string `json:"client_secret"` // client_secret
	Cache        cache.Cache
}
