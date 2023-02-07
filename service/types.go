package service

import (
	"github.com/gofiber/fiber/v2"
)

type Service struct {
	FiberApp *fiber.App
}

var s Service
