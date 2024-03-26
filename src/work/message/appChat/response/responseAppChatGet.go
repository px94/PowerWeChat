package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseAppChatGet struct {
	response.ResponseWork

	ChatInfo *power.HashMap `json:"chat_info"`
}
