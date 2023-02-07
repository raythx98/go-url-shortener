package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"log"

	"github.com/raythx98/go-url-shortener/service/command"
	"github.com/raythx98/go-url-shortener/service/invoker"
	"github.com/raythx98/go-url-shortener/service/receiver"
)

func AddLink(c *fiber.Ctx) error {
	log.Println("Adding link...")

	addLinkCommand := &command.AddLinkCommand{Database: receiver.DbInstance}

	if err := c.BodyParser(&addLinkCommand); err != nil {
		log.Println(err)
		return c.Status(400).SendString("Cannot unmarshal payload")
	}

	addLinkInvoker := &invoker.Invoker{
		Command: addLinkCommand,
	}
	shortenedUrl := addLinkInvoker.Invoke()

	addLinkCommandJSON, err := json.Marshal(addLinkCommand)
	if err == nil {
		log.Println(string(addLinkCommandJSON))
	}

	return c.Status(201).Send([]byte(shortenedUrl))
}

func GetFullLink(c *fiber.Ctx) error {
	log.Println("Redirecting...")

	log.Println("From:", c.Params("shortUrl"))
	getFullLinkCommand := &command.GetFullLinkCommand{
		Database:      receiver.DbInstance,
		ShortenedLink: c.Params("shortUrl"),
	}
	getFullLinkInvoker := &invoker.Invoker{
		Command: getFullLinkCommand,
	}
	fullUrl := getFullLinkInvoker.Invoke()
	log.Println("To:", fullUrl)
	return c.Redirect(fullUrl)
}
