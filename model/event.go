package model

import (
	"fmt"
	"io"
)

// Message is the event sent or received from a Channel
type Message struct {
	data []byte
}

// NewMessage creates a new Message
func NewMessage(d []byte) Message {
	return Message{data: d}
}

// Data returns the message data
func (e *Message) Data() []byte {
	return e.data
}

// Write let the message write itself to the given io.Writer
func (e *Message) Write(w io.Writer) {
	fmt.Fprintf(w, "data: %s\n\n", string(e.Data()))
}
