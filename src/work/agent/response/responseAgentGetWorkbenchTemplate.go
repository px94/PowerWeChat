package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
	"github.com/px94/PowerWeChat/v3/src/work/agent/request"
)

type ResponseAgentGetWorkbenchTemplate struct {
	response.ResponseWork

	TemplateType    string                 `json:"type"`
	Image           request.WorkBenchImage `json:"image"`
	ReplaceUserData bool                   `json:"replace_user_data"`
}
