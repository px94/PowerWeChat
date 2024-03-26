package response

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseUserExportGetResult struct {
	response.ResponseWork

	Status   int               `json:"status"`
	DataList []*object.HashMap `json:"data_list"`
}
