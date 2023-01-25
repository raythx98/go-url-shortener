package main

import (
	"github.com/raythx98/go-url-shortener/service"
)

func main() {
	service.Init()
	service.AddHandlers()
}
