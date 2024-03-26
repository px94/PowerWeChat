package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseDataCubeSummary struct {
	response.ResponseMiniProgram

	List []*power.HashMap `json:"list"`
}
