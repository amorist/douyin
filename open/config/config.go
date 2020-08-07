package config

import "github.com/amorist/douyin/open/cache"

// Config config for 抖音开放平台
type Config struct {
	ClientKey    string
	ClientSecret string
	RedirectURL  string
	Cache        cache.Cache
}
