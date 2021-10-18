package client

import (
	"context"
	"log"

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

// GetEvents allows to read the events for the given id uuid.UUID
func (c *RedisClient) GetEvents(ctx context.Context, id uuid.UUID) (chan model.Event, error) {
	pubsub := c.rds.Subscribe(ctx, id.String())
	log.Printf("subscribed %+v\n", pubsub)
	ch := pubsub.Channel()

	events := make(chan model.Event)
	go func() {
		for {
			log.Println("waiting messages")
			select {
			case <-ctx.Done():
				log.Println("closing pubsub")
				pubsub.Close()
				return
			case msg := <-ch:
				log.Printf("message received %+v", msg)
				events <- model.NewEvent([]byte(msg.Payload))
			}
		}
	}()

	return events, nil
}

// Send sends an event to the given id uuid.UUID
func (c *RedisClient) Send(ctx context.Context, ev model.Event, id uuid.UUID) error {
	log.Printf("sending message to redis %+v", ev)
	return c.rds.Publish(ctx, id.String(), ev.Data()).Err()
}
