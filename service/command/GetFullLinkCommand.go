package command

import (
	"context"
	"log"

	"github.com/raythx98/go-url-shortener/service/receiver"
)

type GetFullLinkCommand struct {
	Database      receiver.Database
	ShortenedLink string
}

func (c *GetFullLinkCommand) Execute() string {
	// business logic
	fullLink, err := c.Database.GetFullLink(context.Background(), c.ShortenedLink)
	if err != nil {
		log.Println(err)
		return ""
	}
	return fullLink
}
