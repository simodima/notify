package client

import (
	"context"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"

	"github.com/toretto460/notify/model"
)

// RedisClient is the notification client redis implementation
type RedisClient struct {
	rds *redis.Client
}

// NewRedis creates a new RedisClient
func NewRedis(cli *redis.Client) RedisClient {
	return RedisClient{
		rds: cli,
	}
}

// Init will initialize the client
func (c *RedisClient) Init(context.Context, uuid.UUID) error {
	return nil
}

// GetEvents allows to read the events for the given id uuid.UUID
func (c *RedisClient) GetEvents(ctx context.Context, id uuid.UUID) (chan model.Message, error) {
	pubsub := c.rds.Subscribe(ctx, id.String())
	ch := pubsub.Channel()

	events := make(chan model.Message)
	go func() {
		for {
			select {
			case <-ctx.Done():
				pubsub.Close()
				return
			case msg := <-ch:
				events <- model.NewMessage([]byte(msg.Payload))
			}
		}
	}()

	return events, nil
}

// Send sends an event to the given id uuid.UUID
func (c *RedisClient) Send(ctx context.Context, ev model.Message, id uuid.UUID) error {
	return c.rds.Publish(ctx, id.String(), ev.Data()).Err()
}
