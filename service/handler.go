package service

import (
	"fmt"
	"net/http"

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
	fmt.Println("Turning on...")
	onCommand := &command.OnCommand{
		Device: &receiver.Tv{},
	}
	onButton := &invoker.Button{
		Command: onCommand,
	}
	onButton.Press()
	fmt.Println("Turned on...")
	return &url_shortener_proto.TurnOnResponse{}, nil
}

func (s *Server) TurnOff(ctx context.Context, request *url_shortener_proto.TurnOffRequest) (*url_shortener_proto.TurnOffResponse, error) {
	fmt.Println("Turning off...")
	offCommand := &command.OffCommand{
		Device: &receiver.Tv{},
	}
	offButton := &invoker.Button{
		Command: offCommand,
	}
	offButton.Press()
	fmt.Println("Turned off...")
	return &url_shortener_proto.TurnOffResponse{}, nil
}

func (on OnHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Turning on...\n"))
	onCommand := &command.OnCommand{
		Device: &receiver.Tv{},
	}
	onButton := &invoker.Button{
		Command: onCommand,
	}
	onButton.Press()
	w.Write([]byte("Turned on..."))
}

func (off OffHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Turning off...\n"))
	offCommand := &command.OffCommand{
		Device: &receiver.Tv{},
	}
	offButton := &invoker.Button{
		Command: offCommand,
	}
	offButton.Press()
	w.Write([]byte("Turned off..."))
}
