package handler

import (
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/raythx98/go-url-shortener/service/command"
	"github.com/raythx98/go-url-shortener/service/invoker"
	"github.com/raythx98/go-url-shortener/service/receiver"
)

func TurnOn(c *fiber.Ctx) error {
	log.Println("Turning off...")
	onCommand := &command.OnCommand{
		Device: &receiver.Tv{},
	}
	onButton := &invoker.Button{
		Command: onCommand,
	}
	onButton.Press()
	return c.Status(200).Send([]byte("Turned On..."))
}

func TurnOff(c *fiber.Ctx) error {
	log.Println("Turning off...")
	offCommand := &command.OffCommand{
		Device: &receiver.Tv{},
	}
	offButton := &invoker.Button{
		Command: offCommand,
	}
	offButton.Press()
	return c.Status(200).Send([]byte("Turned Off..."))
}
