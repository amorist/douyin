package util

// CommonError 抖音返回的通用错误json
type CommonError struct {
	ErrCode int64  `json:"error_code"`
	ErrMsg  string `json:"description"`
}
