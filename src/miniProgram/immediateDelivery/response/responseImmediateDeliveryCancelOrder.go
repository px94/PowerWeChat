package response

import "github.com/px94/PowerWeChat/v3/src/kernel/response"

type ResponseImmediateDeliveryCancelOrder struct {
	response.ResponseMiniProgram

	DeductFee int64  `json:"deduct_fee"`
	Desc      string `json:"desc"`
}
