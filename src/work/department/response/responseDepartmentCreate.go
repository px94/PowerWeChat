package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseDepartmentCreate struct {
	response.ResponseWork
	ID int `json:"id"`
}
