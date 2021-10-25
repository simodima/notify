// Copyright © 2021 Simone Di Maulo <toretto460@gmail.com>.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

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
