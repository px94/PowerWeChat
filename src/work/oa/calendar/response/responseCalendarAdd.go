package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseCalendarAdd struct {
	response.ResponseWork

	CalID string `json:"cal_id"`
}
