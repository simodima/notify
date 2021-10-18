package model

import (
	"fmt"
	"io"
)

// Event is the event sent or received from a Channel
type Event struct {
	data []byte
}

// NewEvent creates a new Event
func NewEvent(d []byte) Event {
	return Event{data: d}
}

// Data returns the event data
func (e *Event) Data() []byte {
	return e.data
}

// Write let the event write itself to the given io.Writer
func (e *Event) Write(w io.Writer) {
	fmt.Fprintf(w, "data: %s\n\n", string(e.Data()))
}
