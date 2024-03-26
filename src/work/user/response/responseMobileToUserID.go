package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseMobileToUserID struct {
	response.ResponseWork

	UserID string `json:"userid"`
}

type ResponseConvertToUserID struct {
	response.ResponseWork

	UserID string `json:"userid"`
}
