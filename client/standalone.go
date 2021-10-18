package client

import (
	"context"

	"github.com/google/uuid"

	"github.com/toretto460/notify/model"
)

type StandaloneClient struct {
	events       chan model.Event
	closeContext context.Context
}

func NewStandalone(ctx context.Context) StandaloneClient {
	panic("to be implemented")
}

func (c *StandaloneClient) GetEvents(ctx context.Context, id uuid.UUID) (chan model.Event, error) {
	panic("to be implemented")
}

func (c *StandaloneClient) Send(ctx context.Context, ev model.Event, id uuid.UUID) error {
	panic("to be implemented")
}
