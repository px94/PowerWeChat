package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseServicerDel struct {
	response.ResponseWork

	ResultList []*power.HashMap `json:"result_list"`
}
