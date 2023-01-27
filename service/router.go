package service

import (
	"github.com/raythx98/go-url-shortener/service/handler"
)

func (app *FiberApp) Listen(addr string) error {
	return app.fiberApp.Listen(addr)
}

func (app *FiberApp) RegisterRouter() {
	app.fiberApp.Get("/on", handler.TurnOn)
	app.fiberApp.Get("/off", handler.TurnOff)
}
