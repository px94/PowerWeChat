package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseMsgAuditGetRoom struct {
	response.ResponseWork

	Roomname       string           `json:"roomname"`
	Creator        string           `json:"creator"`
	RoomCreateTime int              `json:"room_create_time"`
	Notice         string           `json:"notice"`
	Members        []*power.HashMap `json:"members"`
}
