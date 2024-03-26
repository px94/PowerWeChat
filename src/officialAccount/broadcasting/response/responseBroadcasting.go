package response

import "github.com/px94/PowerWeChat/v3/src/kernel/response"

type ResponseBroadcastingSend struct {
	response.ResponseOfficialAccount

	MsgID     int `json:"msg_id"`
	MsgDataID int `json:"msg_data_id"`
}

// ---------------------------------------------

type ResponseBroadcastingStatus struct {
	MsgID     int    `json:"msg_id"`
	MsgStatus string `json:"msg_status"`
}
