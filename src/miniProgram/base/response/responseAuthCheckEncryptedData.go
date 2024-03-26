package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseAuthCheckEncryptedData struct {
	response.ResponseMiniProgram

	Valid      bool    `json:"vaild"`
	CreateTime float64 `json:"create_time"`
}
