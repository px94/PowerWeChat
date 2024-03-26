package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveFileRename struct {
	response.ResponseWork

	File *power.HashMap `json:"file"`
}
