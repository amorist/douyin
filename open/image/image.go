package image

import (
	"github.com/amorist/douyin/open/context"
)

const (
	// 上传图片
	imageUploadURL string = "https://open.douyin.com/image/upload?access_token=%s&open_id=%s"
	// 发布图片
	imageCreateURL string = "https://open.douyin.com/image/create?access_token=%s&open_id=%s"
)

// Image 图片
type Image struct {
	*context.Context
}

// Upload 图片上传
func (image *Image) Upload() (err error) {
	return
}
