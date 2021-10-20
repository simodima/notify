package channel

import (
	"context"

	"github.com/google/uuid"

	"github.com/toretto460/notify/model"
)

// FromID creates a channel from the the given id
func FromID(chID string, client cli) (Channel, error) {
	client.Init(context.Background(), chID)

	return Channel{
		id: chID,
		c:  client,
	}, nil
}

// NewChannel creates a new Channel
func NewChannel(client cli) (Channel, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return Channel{}, err
	}

	client.Init(context.Background(), id.String())

	return Channel{
		id: id.String(),
		c:  client,
	}, nil
}

// Channel represents the channel where the messages can be sent and received
type Channel struct {
	id string
	c  cli
}

// ID returns the channel identifier
func (c *Channel) ID() string {
	return c.id
}

// Send sends a message to the channel
func (c *Channel) Send(ctx context.Context, ev model.Message) error {
	return c.c.Send(ctx, ev, c.id)
}

// Receive returns a go chan to receive messages over the channel
func (c *Channel) Receive(ctx context.Context) (chan model.Message, error) {
	return c.c.GetEvents(ctx, c.id)
}
