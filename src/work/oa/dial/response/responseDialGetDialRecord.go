package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseDialGetDialRecord struct {
	response.ResponseWork

	MeetingidList []*power.HashMap `json:"record"`
}
