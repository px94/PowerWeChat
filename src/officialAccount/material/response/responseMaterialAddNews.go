package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseMaterialAddNews struct {
	response.ResponseOfficialAccount

	MediaID string `json:"media_id"`
}
