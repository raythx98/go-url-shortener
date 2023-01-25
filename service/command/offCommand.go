package command

import (
	// stdlib

	// third party

	// internal github

	"github.com/raythx98/go-url-shortener/service/receiver"
)

type OffCommand struct {
	Device receiver.Device
}

func (c *OffCommand) Execute() {
	c.Device.Off()
	// business logic
}
