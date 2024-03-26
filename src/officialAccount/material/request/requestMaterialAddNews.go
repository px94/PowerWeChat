package request

import "github.com/px94/PowerWeChat/v3/src/kernel/power"

type RequestMaterialAddNews struct {
	Articles []*power.HashMap `json:"articles"`
}
