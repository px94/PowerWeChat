package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseTagGetCorpTagList struct {
	response.ResponseWork

	TagGroups []*CorpTagGroup `json:"tag_group"`
}
