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

package channel

import (
	"context"

	"github.com/toretto460/notify/model"
)

type cli interface {
	Init(context.Context, string) error
	GetEvents(context.Context, string) (chan model.Message, error)
	Send(context.Context, model.Message, string) error
}

// NewFactory creates a new Factory for channels
func NewFactory(client cli) Factory {
	return Factory{cli: client}
}

// Factory is the channel factory.
type Factory struct {
	cli cli
}

// New creates a new channel
func (c *Factory) New() (Channel, error) {
	return NewChannel(c.cli)
}

// Get creates a new channel for the given id
func (c *Factory) Get(id string) (Channel, error) {
	return FromID(id, c.cli)
}
