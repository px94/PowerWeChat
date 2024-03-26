package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseGroupChatOpenGIDToChatID struct {
	response.ResponseWork
	ChatID *power.HashMap `json:"chat_id"`
}
