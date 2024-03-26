package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInRules struct {
	response.ResponseWork
	Info []*power.HashMap `json:"info"`
}
