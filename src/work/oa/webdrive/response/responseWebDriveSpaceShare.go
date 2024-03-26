package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseWebDriveSpaceShare struct {
	response.ResponseWork

	SpaceShareURL string `json:"space_share_url"`
}
