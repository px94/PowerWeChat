package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseMsgAuditGetAgreeInfo struct {
	response.ResponseWork
	AgreeInfo []*power.HashMap `json:"agreeinfo"`
}
