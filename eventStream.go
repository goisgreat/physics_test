package main

import (
	physics "github.com/goisgreat/go-physics"

	tools "github.com/goisgreat/ebiten-tools"
)

// define an event stream object
var eventStream physics.EventStream

func doEventStream() {
	// set event stream
	eventStream = physics.CreateEventStream()
	// push keyboard input onto eventStream.Input (chan byte)
	go tools.PushInputBytes(eventStream.Input)
}
