package response

import "github.com/px94/PowerWeChat/v3/src/kernel/response"

type ResponseAddContactWay struct {
	response.ResponseWork

	ConfigID string `json:"config_id"`
	QRCode   string `json:"qr_code"`
}
