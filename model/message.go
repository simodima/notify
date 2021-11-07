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
	"encoding/json"
)

// Message is the event sent or received from a Channel
type Message struct {
	name string
	data []byte
}

// NewMessage creates a new Message
func NewMessage(name string, d []byte) Message {
	return Message{data: d, name: name}
}

// UnmarshalJSON unmarshal byte slice into a model.Message
func (e *Message) UnmarshalJSON(data []byte) error {
	rawMsg := struct {
		Name string
		Data []byte
	}{}
	err := json.Unmarshal(data, &rawMsg)
	if err != nil {
		return err
	}

	e.data = rawMsg.Data
	e.name = rawMsg.Name

	return nil
}

// MarshalJSON marshal a model.Message into an byte slice
func (e *Message) MarshalJSON() ([]byte, error) {
	rawMsg := struct {
		Name string
		Data []byte
	}{
		Data: e.data,
		Name: e.name,
	}

	return json.Marshal(&rawMsg)
}

// Name returns the message name
func (e *Message) Name() string {
	return e.name
}

// Data returns the message data
func (e *Message) Data() []byte {
	return e.data
}
