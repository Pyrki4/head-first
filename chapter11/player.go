package main

import (
	"golang/chapter11/gadget"
)

func TryOut(player gadget.Player) {
	player.Play("Test Track")
	player.Stop()
	recorder, ok := player.(gadget.TapeRecorder)
	if ok {
		recorder.Record()
	}
}

func main() {
	TryOut(gadget.TapePlayer{})
	TryOut(gadget.TapeRecorder{})
}
