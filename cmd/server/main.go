package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/raythx98/go-url-shortener/service"
	"log"
)

func main() {
	service.Init()

	app := &service.Service{
		FiberApp: fiber.New(),
	}
	app.RegisterRouter()

	log.Println("Listening...")
	app.Listen(":9281")

	// stdlib http implementation
	/*
		on := OnHandler{}
		off := OffHandler{}

		muxHttp := http.NewServeMux()
			// Use the mux.Handle() function to register this with our new servemux,
			// so it acts as the customhandler for all incoming requests with the URL path /test.
			muxHttp.Handle("/on", on)
			muxHttp.Handle("/off", off)

			// Then we create a new server and start listening for incoming requests
			// with the http.ListenAndServe() function, passing in our servemux for it to
			// match requests against as the second parameter.
			err = http.ListenAndServe(":9281", muxHttp)
			if err != nil {
				log.Fatalf("failed to listen to HTTP port 9281: %v", err)
			}
	*/
}
