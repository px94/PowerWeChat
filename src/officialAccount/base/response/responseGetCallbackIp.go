package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetCallBackIP struct {
	response.ResponseOfficialAccount

	IPList []string `json:"ip_list"`
}
