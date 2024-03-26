package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseOperationGetVersionList struct {
	response.ResponseMiniProgram
	CVList []*power.HashMap `json:"cvlist"`
}
