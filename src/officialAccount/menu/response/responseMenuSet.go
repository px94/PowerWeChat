package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseMenuCreate struct {
	response.ResponseOfficialAccount
}

type ResponseMenuCreateConditional struct {
	response.ResponseOfficialAccount

	MenuID string `json:"menuid"`
}
