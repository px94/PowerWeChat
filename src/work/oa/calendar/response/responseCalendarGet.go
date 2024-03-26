package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseCalendarGet struct {
	response.ResponseWork

	CalendarList []*power.HashMap `json:"calendar_list"`
}
