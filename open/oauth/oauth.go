package oauth

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/amorist/douyin/open/context"
	"github.com/amorist/douyin/open/credential"
	"github.com/amorist/douyin/util"
)

const (
	redirectOauthURL      string = "https://open.douyin.com/platform/oauth/connect?client_key=%s&response_type=code&scope=%s&redirect_uri=%s&state=%s"
	accessTokenURL        string = "https://open.douyin.com/oauth/access_token?client_key=%s&client_secret=%s&code=%s&grant_type=authorization_code"
	refreshAccessTokenURL string = "https://open.douyin.com/oauth/oauth/refresh_token?client_key=%s&grant_type=refresh_token&refresh_token=%s"
	clientTokenURL        string = "https://open.douyin.com/oauth/oauth/client_token/?client_key=%s&client_secret=%s&grant_type=client_credential"
	userInfoURL           string = "https://open.douyin.com/oauth/oauth/userinfo?access_token=%s&open_id=%s"
)

// Oauth 保存用户授权信息
type Oauth struct {
	*context.Context
}

// NewOauth 实例化授权信息
func NewOauth(context *context.Context) *Oauth {
	auth := new(Oauth)
	auth.Context = context
	return auth
}

// GetRedirectURL 获取授权码的url地址
func (oauth *Oauth) GetRedirectURL(state string) string {
	uri := url.QueryEscape(oauth.RedirectURL)
	return fmt.Sprintf(redirectOauthURL, oauth.ClientKey, oauth.Scopes, uri, state)
}

type accessTokenRes struct {
	Message string                 `json:"message"`
	Data    credential.AccessToken `json:"data"`
}

// GetUserAccessToken 通过网页授权的code 换取access_token
func (oauth *Oauth) GetUserAccessToken(code string) (accessToken credential.AccessToken, err error) {
	uri := fmt.Sprintf(accessTokenURL, oauth.ClientKey, oauth.ClientSecret, code)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result accessTokenRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}

	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("GetUserAccessToken error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}

	err = oauth.SetAccessToken(&result.Data)
	if err != nil {
		return
	}

	return
}

// UserInfo 用户信息
type UserInfo struct {
	util.CommonError

	Avatar       string `json:"avatar"`
	City         string `json:"city"`
	Country      string `json:"country"`
	EAccountRole string `json:"e_account_role"`
	Gender       int32  `json:"gender"`
	Nickname     string `json:"nickname"`
	OpenID       string `json:"open_id"`
	Province     string `json:"province"`
	Unionid      string `json:"union_id"`
}

type userInforRes struct {
	Message string   `json:"message"`
	Data    UserInfo `json:"data"`
}

// GetUserInfo 获取用户信息.
func (oauth *Oauth) GetUserInfo(accessToken, openID string) (userInfo *UserInfo, err error) {
	uri := fmt.Sprintf(userInfoURL, accessToken, openID)
	var response []byte
	response, err = util.HTTPGet(uri)
	if err != nil {
		return
	}
	var result userInforRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("GetUserInfo error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	userInfo = &result.Data
	return
}
