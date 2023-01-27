package handler

import (
	"log"

	"golang.org/x/net/context"

	"github.com/raythx98/go-url-shortener/pb"
	"github.com/raythx98/go-url-shortener/service/command"
	"github.com/raythx98/go-url-shortener/service/invoker"
	"github.com/raythx98/go-url-shortener/service/receiver"
)

type Server struct {
	url_shortener_proto.UnimplementedUrlShortenerServiceServer
}

func (s *Server) TurnOn(ctx context.Context, request *url_shortener_proto.TurnOnRequest) (*url_shortener_proto.TurnOnResponse, error) {
	log.Println("Turning on...")
	onCommand := &command.OnCommand{
		Device: &receiver.Tv{},
	}
	onButton := &invoker.Button{
		Command: onCommand,
	}
	onButton.Press()
	return &url_shortener_proto.TurnOnResponse{Response: "Turned on..."}, nil
}

func (s *Server) TurnOff(ctx context.Context, request *url_shortener_proto.TurnOffRequest) (*url_shortener_proto.TurnOffResponse, error) {
	log.Println("Turning off...")
	offCommand := &command.OffCommand{
		Device: &receiver.Tv{},
	}
	offButton := &invoker.Button{
		Command: offCommand,
	}
	offButton.Press()
	return &url_shortener_proto.TurnOffResponse{Response: "Turned off..."}, nil
}
