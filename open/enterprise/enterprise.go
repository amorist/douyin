package enterprise

import (
	"github.com/amorist/douyin/open/context"
	"github.com/amorist/douyin/open/enterprise/groupon"
	"github.com/amorist/douyin/open/enterprise/im"
	"github.com/amorist/douyin/open/enterprise/interactive"
	"github.com/amorist/douyin/open/enterprise/user"
)

// Enterprise .
type Enterprise struct {
	*context.Context
}

// NewEnterprise .
func NewEnterprise(context *context.Context) *Enterprise {
	enterprise := new(Enterprise)
	enterprise.Context = context
	return enterprise
}

// GetUser 粉丝管理接口.
func (enterprise *Enterprise) GetUser() *user.User {
	return user.NewUser(enterprise.Context)
}

// GetGroupon 团购活动管理接口.
func (enterprise *Enterprise) GetGroupon() *groupon.Groupon {
	return groupon.NewGroupon(enterprise.Context)
}

// GetIM 消息管理接口.
func (enterprise *Enterprise) GetIM() *im.IM {
	return im.NewIM(enterprise.Context)
}

// GetInteractive 互动管理接口.
func (enterprise *Enterprise) GetInteractive() *interactive.Interactive {
	return interactive.NewInteractive(enterprise.Context)
}
