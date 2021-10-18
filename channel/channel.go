package channel

import (
	"context"

	"github.com/google/uuid"

	"github.com/toretto460/notify/model"
)

type cli interface {
	GetEvents(context.Context, uuid.UUID) (chan model.Event, error)
	Send(context.Context, model.Event, uuid.UUID) error
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
func (c *Factory) New() (model.Channel, error) {
	return model.NewChannel(c.cli)
}

// Get creates a new channel for the given id
func (c *Factory) Get(id string) (model.Channel, error) {
	return model.FromID(id, c.cli)
}
