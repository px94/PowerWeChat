package response

import (
	response2 "github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseExpressPreviewTemplate struct {
	*response2.ResponseMiniProgram

	WaybillID               string `json:"waybill_id"`
	RenderedWaybillTemplate string `json:"rendered_waybill_template"`
}
