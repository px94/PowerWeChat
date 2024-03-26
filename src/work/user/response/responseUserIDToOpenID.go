package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserIDToOpenID struct {
	response.ResponseWork

	OpenID string `json:"openid"`
}
