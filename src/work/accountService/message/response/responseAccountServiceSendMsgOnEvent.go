package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseAccountServiceSendMsgOnEvent struct {
	response.ResponseWork

	MsgID string `json:"msgid"`
}
