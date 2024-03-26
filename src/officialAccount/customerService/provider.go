package customerService

import (
	"github.com/px94/PowerWeChat/v3/src/kernel"
	"github.com/px94/PowerWeChat/v3/src/officialAccount/customerService/session"
)

func RegisterProvider(app kernel.ApplicationInterface) (*Client, *session.Client, error) {

	client, err := NewClient(app)
	if err != nil {
		return nil, nil, err
	}

	sessionClient, err := session.NewClient(app)
	if err != nil {
		return nil, nil, err
	}

	return client, sessionClient, nil
}
