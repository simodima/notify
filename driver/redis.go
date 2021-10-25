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

package driver

import (
	"context"

	"github.com/go-redis/redis/v8"

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
func (c *RedisClient) Init(context.Context, string) error {
	return nil
}

// GetEvents allows to read the events for the given id
func (c *RedisClient) GetEvents(ctx context.Context, id string) (chan model.Message, error) {
	pubsub := c.rds.Subscribe(ctx, id)
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

// Send sends an event to the given id
func (c *RedisClient) Send(ctx context.Context, ev model.Message, id string) error {
	return c.rds.Publish(ctx, id, ev.Data()).Err()
}
