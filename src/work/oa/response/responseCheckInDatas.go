package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseCheckInDatas struct {
	response.ResponseWork

	Datas []*power.HashMap `json:"datas"`
}
