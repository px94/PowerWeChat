package miniProgram

import "github.com/px94/PowerWeChat/v3/src/kernel"

type Client struct {
	BaseClient *kernel.BaseClient
}

func NewClient(app *kernel.ApplicationInterface) (*Client, error) {
	baseClient, err := kernel.NewBaseClient(app, nil)
	if err != nil {
		return nil, err
	}
	return &Client{
		BaseClient: baseClient,
	}, nil
}
