package invoker

import (
	"github.com/raythx98/go-url-shortener/service/command"
)

type Button struct {
	Command command.Command
}

func (b *Button) Press() {
	/*
		start goroutine to execute command,
		suitable if response is not needed - EDA
	*/
	// go b.Command.Execute()
	b.Command.Execute()
}
