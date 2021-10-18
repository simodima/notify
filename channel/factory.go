package channel

import (
	"context"

	"github.com/google/uuid"

	"github.com/toretto460/notify/model"
)

type cli interface {
	Init(context.Context, uuid.UUID) error
	GetEvents(context.Context, uuid.UUID) (chan model.Message, error)
	Send(context.Context, model.Message, uuid.UUID) error
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
