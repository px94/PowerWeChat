package auth

import (
	"github.com/px94/PowerWeChat/v3/src/kernel"
)

func RegisterProvider(app kernel.ApplicationInterface) (*AccessToken, error) {

	return NewAccessToken(&app)

}

func RegisterAuthProvider(app kernel.ApplicationInterface) (*Client, error) {

	return NewClient(&app)
}
