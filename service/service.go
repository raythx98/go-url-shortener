package service

import (
	// stdlib
	"fmt"
	"log"
	"net"

	// third party
	"github.com/gofiber/fiber/v2"
	"google.golang.org/grpc"

	// internal GitHub

	// internal project
	pb "github.com/raythx98/go-url-shortener/pb"
	"github.com/raythx98/go-url-shortener/service/handler"
)

func StartService() {
	log.Println("Listening...")

	// can listen from any IP Address
	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", 9282))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterUrlShortenerServiceServer(grpcServer, &handler.Server{})
	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatalf("failed to listen to gRPC port 9282: %v", err)
		}
	}()

	app := &FiberApp{
		fiberApp: fiber.New(),
	}
	app.RegisterRouter()
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
