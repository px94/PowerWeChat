package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveFileUpload struct {
	response.ResponseWork

	FileID string `json:"fileid"`
}
