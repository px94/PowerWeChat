package response

import "github.com/px94/PowerWeChat/v3/src/kernel/response"

type ResponseGetFollowUserList struct {
	response.ResponseWork

	FollowUser []string `json:"follow_user"` // ["zhangsan","tagid2"]
}
