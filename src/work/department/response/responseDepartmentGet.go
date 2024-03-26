package response

import (
	"github.com/px94/PowerWeChat/v3/src/kernel/models"
	"github.com/px94/PowerWeChat/v3/src/kernel/response"
)

type ResponseDepartmentGet struct {
	response.ResponseWork
	Department *models.Department `json:"department"`
}
