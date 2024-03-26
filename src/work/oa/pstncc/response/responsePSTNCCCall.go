package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponsePSTNCCCall struct {
	response.ResponseWork

	States []*power.HashMap `json:"states"`
}
