package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseIMGSuperResolution struct {
	response.ResponseMiniProgram
	MediaID string `json:"media_id"`
}
