package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseAuthGetPaidUnionID struct {
	UnionID string `json:"unionid"`

	response.ResponseMiniProgram
}
