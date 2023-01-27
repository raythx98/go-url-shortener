package receiver

import (
	"log"
)

type Tv struct {
	isRunning bool
}

func (t *Tv) On() {
	t.isRunning = true
	log.Println("Success: on")
}

func (t *Tv) Off() {
	t.isRunning = false
	log.Println("Success: off")
}
