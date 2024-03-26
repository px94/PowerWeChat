package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	response2 "github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseExpressGetContact struct {
	*response2.ResponseMiniProgram

	WayBillID int              `json:"waybill_id"`
	Sender    []*power.HashMap `json:"sender"`
	Receiver  []*power.HashMap `json:"receiver"`
}
