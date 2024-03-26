package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveFileMove struct {
	response.ResponseWork

	FileList *power.HashMap `json:"file_list"`
}
