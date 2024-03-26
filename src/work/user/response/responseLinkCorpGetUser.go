package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseLinkCorpGetUser struct {
	response.ResponseWork

	UserInfo *power.HashMap `json:"user_info"`
}
