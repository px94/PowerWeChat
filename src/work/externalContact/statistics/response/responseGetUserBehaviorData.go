package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetUserBehaviorData struct {
	response.ResponseWork

	MomentList []*power.HashMap `json:"behavior_data"`
}
