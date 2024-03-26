package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInRecords struct {
	response.ResponseWork

	CheckInData []*power.HashMap `json:"checkindata"`
}
