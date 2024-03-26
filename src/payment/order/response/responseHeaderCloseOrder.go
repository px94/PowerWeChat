package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseHeaderCloseOrdr struct {
	response.ResponsePayment

	Status string `header:"status"`
}
