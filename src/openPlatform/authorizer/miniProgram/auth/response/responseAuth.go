package response

import "github.com/px94/PowerWeChat/v3/src/kernel/response"

type ResponseSession struct {
	response.ResponseOpenPlatform

	OpenID     string `json:"openid"`
	SessionKey string `json:"session_key"`
	UnionID    string `json:"unionid"`
}
