package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
	"net/http"

	pb "github.com/raythx98/go-url-shortener/pb"
)

func main() {
	callHttp()
	callGrpc()
}

func callGrpc() {
	log.Println("Calling GRPC: ")
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9282", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	c := pb.NewUrlShortenerServiceClient(conn)

	response, err := c.TurnOn(context.Background(), &pb.TurnOnRequest{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", response)

	offResponse, err := c.TurnOff(context.Background(), &pb.TurnOffRequest{})
	if err != nil {
		log.Fatalf("Error when calling SayHello: %s", err)
	}
	log.Printf("Response from server: %s", offResponse)
}

func callHttp() {
	log.Println("Calling HTTP: ")
	resp, err := http.Get("http://localhost:9281/on")
	if err != nil {
		log.Fatalln(err)
	}
	text, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Response from server: %s", text)
	offResp, err := http.Get("http://localhost:9281/off")
	if err != nil {
		log.Fatalln(err)
	}
	text, err = io.ReadAll(offResp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Response from server: %s", text)
}
