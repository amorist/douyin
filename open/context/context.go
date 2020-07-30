package context

import (
	"github.com/amorist/douyin/open/config"
	"github.com/amorist/douyin/open/credential"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
