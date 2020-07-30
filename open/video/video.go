package video

import "github.com/amorist/douyin/open/context"

const (
	// 上传视频
	videoUploadURL string = "https://open.douyin.com/video/upload?access_token=%s&open_id=%s"
	// 分片初始化上传
	videoPartInitURL string = "https://open.douyin.com/video/part/init?access_token=%s&open_id=%s"
	// 分片上传
	videoPartUploadURL string = "https://open.douyin.com/video/part/upload?access_token=%s&open_id=%s"
	// 分片完成上传
	videoPartCompleteURL string = "https://open.douyin.com/video/part/complete?access_token=%s&open_id=%s"
	// 创建视频
	videoCreateURL string = "https://open.douyin.com/video/create?access_token=%s&open_id=%s"
)

// Video 视频
type Video struct {
	*context.Context
}
