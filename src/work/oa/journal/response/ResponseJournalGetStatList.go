package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/power"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseJournalGetStatList struct {
	response.ResponseWork

	StatList *power.HashMap `json:"stat_list"`
}
