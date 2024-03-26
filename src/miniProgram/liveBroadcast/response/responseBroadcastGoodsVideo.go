package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseBroadcastGoodsVideo struct {
	response.ResponseMiniProgram

	URL int `json:"url"`
}
