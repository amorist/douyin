package open

import (
	"douyin/open/config"
	"douyin/open/context"
	"douyin/open/credential"
	"douyin/open/enterprise"
	externaldata "douyin/open/external_data"
	"douyin/open/image"
	"douyin/open/oauth"
	"douyin/open/pay"
	"douyin/open/poi"
	"douyin/open/search"
	"douyin/open/user"
	"douyin/open/video"
)

//API 抖音开放平台API
type API struct {
	ctx *context.Context
}

//NewOpenAPI 实例抖音开放平台API
func NewOpenAPI(cfg *config.Config) *API {
	defaultAkHandle := credential.NewDefaultAccessToken(cfg.ClientKey, cfg.ClientSecret, credential.CacheKeyPrefix, cfg.Cache)
	ctx := &context.Context{
		Config:            cfg,
		AccessTokenHandle: defaultAkHandle,
	}
	return &API{ctx: ctx}
}

// SetAccessTokenHandle 自定义access_token获取方式
func (api *API) SetAccessTokenHandle(accessTokenHandle credential.AccessTokenHandle) {
	api.ctx.AccessTokenHandle = accessTokenHandle
}

// GetContext get Context
func (api *API) GetContext() *context.Context {
	return api.ctx
}

// GetAccessToken 获取access_token
func (api *API) GetAccessToken(openID string) (string, error) {
	return api.ctx.GetAccessToken(openID)
}

// GetClientToken 获取client_token
func (api *API) GetClientToken() (string, error) {
	clientToken, err := api.ctx.GetClientToken()
	if err != nil {
		return "", err
	}
	return clientToken.AccessToken, nil
}

// GetOauth oauth2网页授权
func (api *API) GetOauth() *oauth.Oauth {
	return oauth.NewOauth(api.ctx)
}

// GetVideo 视频管理接口
func (api *API) GetVideo() *video.Video {
	return video.NewVideo(api.ctx)
}

func (api *API) GetExternalData() *externaldata.ExternalData {
	return externaldata.NewExternalData(api.ctx)
}

// GetImage 图片管理接口
func (api *API) GetImage() *image.Image {
	return image.NewImage(api.ctx)
}

// GetUser 用户管理接口
func (api *API) GetUser() *user.User {
	return user.NewUser(api.ctx)
}

// GetPoi 生活服务管理接口
func (api *API) GetPoi() *poi.Poi {
	return poi.NewPoi(api.ctx)
}

// GetPay 支付管理接口
func (api *API) GetPay() *pay.Pay {
	return pay.NewPay(api.ctx)
}

// GetSearch 搜索管理接口
func (api *API) GetSearch() *search.Search {
	return search.NewSearch(api.ctx)
}

// GetEnterprise 企业管理接口
func (api *API) GetEnterprise() *enterprise.Enterprise {
	return enterprise.NewEnterprise(api.ctx)
}
