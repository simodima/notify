package model

import (
	"context"

	"github.com/google/uuid"
)

type channelClient interface {
	Send(context.Context, Event, uuid.UUID) error
	GetEvents(context.Context, uuid.UUID) (chan Event, error)
}

// FromID creates a channel from the the given id
func FromID(chID string, client channelClient) (Channel, error) {
	id, err := uuid.Parse(chID)
	if err != nil {
		return Channel{}, err
	}

	return Channel{
		id: id,
		c:  client,
	}, nil
}

// NewChannel creates a new Channel
func NewChannel(client channelClient) (Channel, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Channel{}, err
	}

	return Channel{
		id: id,
		c:  client,
	}, nil
}

// Channel represents the channel where the e
type Channel struct {
	id uuid.UUID
	c  channelClient
}

// ID the channel uuid
func (c *Channel) ID() uuid.UUID {
	return c.id
}

// Send sends a message to the channel
func (c *Channel) Send(ctx context.Context, ev Event) error {
	return c.c.Send(ctx, ev, c.id)
}

// Receive returns a go chan to receive messages over the channel
func (c *Channel) Receive(ctx context.Context) (chan Event, error) {
	return c.c.GetEvents(ctx, c.id)
}
