package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingGetLivingCode struct {
	response.ResponseWork

	LivingCode int `json:"living_code"`
}
