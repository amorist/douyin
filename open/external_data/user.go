package externaldata

import (
	"douyin/util"
	"encoding/json"
	"fmt"
)

type QueryURL string

const (
	queryFix            = "open_id=%s&access_token=%s&date_type=%d"
	FansURI    QueryURL = "https://open.douyin.com/data/external/user/fans?%s"
	ItemURI    QueryURL = "https://open.douyin.com/data/external/user/item?%s"
	LikeURI    QueryURL = "https://open.douyin.com/data/external/user/like?%s"
	CommentURI QueryURL = "https://open.douyin.com/data/external/user/comment?%s"
	ShareURI   QueryURL = "https://open.douyin.com/data/external/user/share?%s"
	ProfileURI QueryURL = "https://open.douyin.com/data/external/user/profile?%s"
)

func (e *ExternalData) GetProfileData(openid string, dateType int64, queryURL QueryURL) (*Res, error) {
	at, err := e.GetAccessToken(openid)
	if err != nil {
		return nil, err
	}
	fullUrl := e.buildRequest(openid, dateType, at, queryURL)
	res := new(Res)
	err = e.getHandler(fullUrl, res)
	return res, err
}

func (*ExternalData) GetURLMapping() map[string]QueryURL {
	return map[string]QueryURL{"fans": FansURI, "item": ItemURI, "like": LikeURI,
		"comment": CommentURI, "share": ShareURI, "profile": ProfileURI}
}

func (e *ExternalData) buildRequest(openid string, dateType int64, at string, uri QueryURL) string {
	return fmt.Sprintf(string(uri), fmt.Sprintf(queryFix, openid, at, dateType))
}

func (e *ExternalData) getHandler(uri string, result *Res) error {
	var (
		response []byte
		err      error
	)

	response, err = util.HTTPGet(uri)
	if err != nil {
		return err
	}
	err = json.Unmarshal(response, &result)
	if err != nil {
		return err
	}
	if result.DataRes.ErrCode != 0 {
		err = fmt.Errorf("handler error : errcode=%v , errmsg=%v", result.DataRes.ErrCode, result.DataRes.Description)
		return err
	}
	return nil
}
