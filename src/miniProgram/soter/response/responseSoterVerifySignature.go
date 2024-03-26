package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseSoterVerifySignature struct {
	response.ResponseMiniProgram

	IsOK bool `json:"is_ok"`
}
