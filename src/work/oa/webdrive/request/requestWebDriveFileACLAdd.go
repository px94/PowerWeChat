package request

import "github.com/px94/PowerWeChat/v3/src/kernel/power"

type RequestWebDriveFileACLAdd struct {
	UserID   string           `json:"userid"`
	FileID   string           `json:"fileid"`
	AuthInfo []*power.HashMap `json:"auth_info"`
}
