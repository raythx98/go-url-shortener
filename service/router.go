package service

import (
	"github.com/raythx98/go-url-shortener/service/handler"
)

func (s *Service) Listen(addr string) error {
	return s.FiberApp.Listen(addr)
}

func (s *Service) RegisterRouter() {
	s.FiberApp.Post("/", handler.AddLink)
	s.FiberApp.Get("/:shortUrl", handler.GetFullLink)
}
