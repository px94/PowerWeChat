package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseApprovalNoList struct {
	response.ResponseWork
	SpNoList []string `json:"sp_no_list"`
}
