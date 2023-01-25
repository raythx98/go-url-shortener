package service

import (
	"fmt"
	"net/http"

	"google.golang.org/grpc"
)

func AddHandlers() {
	muxHttp := http.NewServeMux()
	muxGrpc := http.NewServeMux()

	on := OnHandler{}
	off := OffHandler{}

	grpcServer := grpc.NewServer()

	// Use the mux.Handle() function to register this with our new servemux,
	// so it acts as the customhandler for all incoming requests with the URL path /test.
	muxHttp.Handle("/on", on)
	muxHttp.Handle("/off", off)

	fmt.Println("Listening...")

	// Then we create a new server and start listening for incoming requests
	// with the http.ListenAndServe() function, passing in our servemux for it to
	// match requests against as the second parameter.
	go http.ListenAndServe(":9281", muxHttp)
	http.ListenAndServe(":9282", muxGrpc)
}
