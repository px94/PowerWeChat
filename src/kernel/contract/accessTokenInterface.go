package contract

import (
	"github.com/ArtisanCloud/PowerLibs/v3/object"
	response2 "github.com/px94/PowerWeChat/v3/src/kernel/response"
	"net/http"
)

type (
	AccessTokenInterface interface {
		GetToken(refresh bool) (resToken *response2.ResponseGetToken, err error)
		Refresh() AccessTokenInterface
		ApplyToRequest(request *http.Request, requestOptions *object.HashMap) (*http.Request, error)
	}
)
