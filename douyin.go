package douyin

import (
	"github.com/amorist/douyin/open/config"
	"github.com/amorist/douyin/open/context"
	"github.com/amorist/douyin/open/credential"
	"github.com/amorist/douyin/open/enterprise"
	"github.com/amorist/douyin/open/image"
	"github.com/amorist/douyin/open/oauth2"
	"github.com/amorist/douyin/open/pay"
	"github.com/amorist/douyin/open/poi"
	"github.com/amorist/douyin/open/search"
	"github.com/amorist/douyin/open/user"
	"github.com/amorist/douyin/open/video"
)

// Douyin .
type Douyin struct {
	ctx *context.Context
}

// NewDouyin .
func NewDouyin(cfg *config.Config) *Douyin {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.ClientKey, cfg.ClientSecret, credential.CacheKeyPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &Douyin{ctx: ctx}
}

// SetAccessTokenHandle 自定义access_token获取方式
func (douyin *Douyin) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	douyin.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (douyin *Douyin) GetContext() *context.Context {
	return douyin.ctx
}

// GetAccessToken 获取access_token
func (douyin *Douyin) GetAccessToken(openID string) (string, error) {
	return douyin.ctx.GetAccessToken(openID)
}

// GetClientToken 获取client_token
func (douyin *Douyin) GetClientToken() (string, error) {
	clientToken, err := douyin.ctx.GetClientToken()
	if err != nil {
		return "", err
	}
	return clientToken.AccessToken, nil
}

// GetOauth oauth2网页授权
func (douyin *Douyin) GetOauth() *oauth2.Oauth {
	return oauth2.NewOauth(douyin.ctx)
}

// GetVideo 视频管理接口
func (douyin *Douyin) GetVideo() *video.Video {
	return video.NewVideo(douyin.ctx)
}

// GetImage 图片管理接口
func (douyin *Douyin) GetImage() *image.Image {
	return image.NewImage(douyin.ctx)
}

// GetUser 用户管理接口
func (douyin *Douyin) GetUser() *user.User {
	return user.NewUser(douyin.ctx)
}

// GetPoi 生活服务管理接口
func (douyin *Douyin) GetPoi() *poi.Poi {
	return poi.NewPoi(douyin.ctx)
}

// GetPay 支付管理接口
func (douyin *Douyin) GetPay() *pay.Pay {
	return pay.NewPay(douyin.ctx)
}

// GetSearch 搜索管理接口
func (douyin *Douyin) GetSearch() *search.Search {
	return search.NewSearch(douyin.ctx)
}

// GetEnterprise 企业管理接口
func (douyin *Douyin) GetEnterprise() *enterprise.Enterprise {
	return enterprise.NewEnterprise(douyin.ctx)
}
