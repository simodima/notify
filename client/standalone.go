package client

import (
	"context"
	"errors"
	"sync"

	"github.com/toretto460/notify/model"
)

type StandaloneClient struct {
	events map[string]chan model.Message
	mu     sync.Mutex
}

func NewStandalone(ctx context.Context) StandaloneClient {
	return StandaloneClient{
		events: make(map[string]chan model.Message),
		mu:     sync.Mutex{},
	}
}

func (c *StandaloneClient) getOrCreateChan(id string) chan model.Message {
	if events, ok := c.events[id]; ok {
		return events
	}
	c.mu.Lock()
	defer c.mu.Unlock()

	c.events[id] = make(chan model.Message)
	return c.events[id]
}

// Init will initialize the client
func (c *StandaloneClient) Init(ctx context.Context, id string) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.events[id] = make(chan model.Message)
	return nil
}

func (c *StandaloneClient) GetEvents(ctx context.Context, id string) (chan model.Message, error) {
	return c.getOrCreateChan(id), nil
}

func (c *StandaloneClient) Send(ctx context.Context, ev model.Message, id string) error {
	if events, ok := c.events[id]; ok {
		events <- ev
		return nil
	}

	return errors.New("channel is not valid")
}
