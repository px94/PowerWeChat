package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseOpenIDToUserID struct {
	response.ResponseWork

	UserID string `json:"userid"`
}
