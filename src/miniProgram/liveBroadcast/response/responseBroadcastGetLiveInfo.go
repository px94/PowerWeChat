package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGetLiveInfo struct {
	response.ResponseMiniProgram

	Total    int              `json:"total"`
	RoomInfo []*power.HashMap `json:"room_info"`
}
