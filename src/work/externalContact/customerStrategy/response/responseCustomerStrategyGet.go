package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerStrategyGet struct {
	response.ResponseWork

	Strategy *power.HashMap `json:"momentStrategy"`
}
