package channel

import (
	"context"

	"github.com/google/uuid"
	"github.com/toretto460/notify/model"
)

// FromID creates a channel from the the given id
func FromID(chID string, client cli) (Channel, error) {
	id, err := uuid.Parse(chID)
	if err != nil {
		return Channel{}, err
	}

	client.Init(context.Background(), id)

	return Channel{
		id: id,
		c:  client,
	}, nil
}

// NewChannel creates a new Channel
func NewChannel(client cli) (Channel, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Channel{}, err
	}

	client.Init(context.Background(), id)

	return Channel{
		id: id,
		c:  client,
	}, nil
}

// Channel represents the channel where the messages can be sent and received
type Channel struct {
	id uuid.UUID
	c  cli
}

// ID returns the channel identifier
func (c *Channel) ID() string {
	return c.id.String()
}

// Send sends a message to the channel
func (c *Channel) Send(ctx context.Context, ev model.Message) error {
	return c.c.Send(ctx, ev, c.id)
}

// Receive returns a go chan to receive messages over the channel
func (c *Channel) Receive(ctx context.Context) (chan model.Message, error) {
	return c.c.GetEvents(ctx, c.id)
}
