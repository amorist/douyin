package douyin

import (
	"github.com/amorist/douyin/open"
	"github.com/amorist/douyin/open/cache"
	"github.com/amorist/douyin/open/config"
)

// Douyin .
type Douyin struct {
	cache cache.Cache
}

// New init.
func New() *Douyin {
	return &Douyin{}
}

//SetCache 设置cache
func (douyin *Douyin) SetCache(cahce cache.Cache) {
	douyin.cache = cahce
}

//GetOpenAPI 获取开放平台API实例
func (douyin *Douyin) GetOpenAPI(cfg *config.Config) *open.API {
	if cfg.Cache == nil {
		cfg.Cache = douyin.cache
	}
	return open.NewOpenAPI(cfg)
}
