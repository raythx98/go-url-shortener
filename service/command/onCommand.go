package command

import (
	"github.com/raythx98/go-url-shortener/service/receiver"
)

type OnCommand struct {
	Device receiver.Device
}

func (c *OnCommand) Execute() {
	c.Device.On()
	// business logic
}
