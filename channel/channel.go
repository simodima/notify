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

	"github.com/google/uuid"

	"github.com/toretto460/notify/model"
)

// FromID creates a channel from the the given id
func FromID(chID string, d driver) (Channel, error) {
	d.Init(context.Background(), chID)

	return Channel{
		id: chID,
		d:  d,
	}, nil
}

// New creates a new Channel
func New(d driver) (Channel, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Channel{}, err
	}

	d.Init(context.Background(), id.String())

	return Channel{
		id: id.String(),
		d:  d,
	}, nil
}

// Channel represents the channel where the messages can be sent and received
type Channel struct {
	id string
	d  driver
}

// ID returns the channel identifier
func (c *Channel) ID() string {
	return c.id
}

// Send sends a message to the channel
func (c *Channel) Send(ctx context.Context, ev model.Message) error {
	return c.d.Send(ctx, ev, c.id)
}

// Receive returns a go chan to receive messages over the channel
func (c *Channel) Receive(ctx context.Context) (chan model.Message, error) {
	return c.d.Receive(ctx, c.id)
}
