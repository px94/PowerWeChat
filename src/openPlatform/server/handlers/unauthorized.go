package handlers

import (
	"github.com/px94/PowerWeChat/v3/src/kernel"
	"github.com/px94/PowerWeChat/v3/src/kernel/contract"
	"net/http"
)

type Unauthorized struct {
	contract.EventHandlerInterface

	App *kernel.ApplicationInterface
}

func NewUnauthorized(app *kernel.ApplicationInterface) *Unauthorized {
	handler := &Unauthorized{
		App: app,
	}
	return handler
}

func (handler *Unauthorized) Handle(request *http.Request, header contract.EventInterface, content interface{}) interface{} {

	return nil
}
