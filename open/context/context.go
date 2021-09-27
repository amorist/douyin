package context

import (
	"douyin/open/config"
	"douyin/open/credential"
)

// Context struct
type Context struct {
	*config.Config
	credential.AccessTokenHandle
}
