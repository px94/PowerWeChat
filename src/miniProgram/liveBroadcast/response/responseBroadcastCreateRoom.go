package response

import "github.com/px94/PowerWeChat/v3/src/kernel/response"

type ResponseBroadcastCreateRoom struct {
	response.ResponseMiniProgram

	RoomID    int    `json:"roomId"`
	QRCodeURL string `json:"qrcode_url"`
}
