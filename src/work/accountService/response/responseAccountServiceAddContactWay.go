package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountServiceAddContactWay struct {
	response.ResponseWork

	URL string `json:"url"`
}
