package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserExportJobs struct {
	response.ResponseWork

	JobID string `json:"jobid"`
}
