package service

import (
	"net/http"

	"github.com/raythx98/go-url-shortener/service/command"
	"github.com/raythx98/go-url-shortener/service/invoker"
	"github.com/raythx98/go-url-shortener/service/receiver"
)

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
