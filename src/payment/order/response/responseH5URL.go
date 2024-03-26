package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseH5URL struct {
	response.ResponsePayment

	H5URL string `json:"h5_url"`
}
