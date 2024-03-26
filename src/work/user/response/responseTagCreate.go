package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagCreate struct {
	response.ResponseWork

	TagID string `json:"tagid"`
}
