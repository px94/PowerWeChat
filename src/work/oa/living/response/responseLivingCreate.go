package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseLivingCreate struct {
	response.ResponseWork

	LivingID int `json:"livingid"`
}
