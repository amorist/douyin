package externaldata

import "douyin/open/context"

// ExternalData 数据开放服务.
type ExternalData struct {
	*context.Context
}

// NewExternalData .
func NewExternalData(context *context.Context) *ExternalData {
	externalData := new(ExternalData)
	externalData.Context = context
	return externalData
}

type Extra struct {
	Now            int64  `json:"now,omitempty"`
	SubDescription string `json:"sub_description,omitempty"`
	SubErrorCode   int64  `json:"sub_error_code,omitempty"`
	Description    string `json:"description,omitempty"`
	ErrorCode      int64  `json:"error_code,omitempty"`
	LogID          string `json:"log_id,omitempty"`
}

type Data struct {
	Description string                   `json:"description"`
	ErrCode     int64                    `json:"err_code"`
	ResultList  []map[string]interface{} `json:"result_list"`
}

type Res struct {
	ExtraRes Extra `json:"extra"`
	DataRes  Data  `json:"data"`
}
