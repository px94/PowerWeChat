package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseGetGroupWelcomeTemplate struct {
	response.ResponseWork

	Text        *power.HashMap `json:"text"`
	Image       *power.HashMap `json:"image"`
	MiniProgram *power.HashMap `json:"miniprogram"`
	File        *power.HashMap `json:"file"`
	Video       *power.HashMap `json:"video"`
}
