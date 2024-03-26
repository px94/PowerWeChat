package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseCustomerStrategyCreate struct {
	response.ResponseWork

	StrategyID *power.HashMap `json:"strategy_id"`
}
