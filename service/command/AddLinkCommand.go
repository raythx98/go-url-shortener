package command

import (
	"context"
	"github.com/aidarkhanov/nanoid"
	database "github.com/raythx98/go-url-shortener/database/sqlc/output"
	"github.com/raythx98/go-url-shortener/service/receiver"
	"log"
)

type AddLinkCommand struct {
	Database     receiver.Database
	FullLink     string `json:"full_link"`
	CustomLink   string `json:"custom_link,omitempty"`
	InvalidateAt string `json:"invalidate_at,omitempty"` // TODO: add datetime unmarshal
	NumRedirects int    `json:"number_redirects,omitempty"`
}

func (c *AddLinkCommand) Execute() string {
	// business logic
	shortenedLink, _ := nanoid.Generate(nanoid.DefaultAlphabet, 8)
	err := c.Database.AddLink(context.Background(), database.AddLinkParams{Link: c.FullLink, ShortenedLink: shortenedLink})
	if err != nil {
		log.Println(err)
	}
	return shortenedLink
}
