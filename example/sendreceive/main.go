package main

import (
	"context"
	"log"
	"time"

	"github.com/toretto460/notify/channel"
	"github.com/toretto460/notify/client"
	"github.com/toretto460/notify/model"
)

func main() {
	ctx := context.TODO()

	client := client.NewStandalone(ctx)
	channeFactory := channel.NewFactory(&client)

	ch, _ := channeFactory.New()

	go func() {
		// Receive messages
		messages, _ := ch.Receive(ctx)
		for msg := range messages {
			log.Printf("Message received [%s]", string(msg.Data()))
		}
	}()

	// Publish messages
	ch.Send(ctx, model.NewMessage([]byte("TEST EVENT 1")))
	ch.Send(ctx, model.NewMessage([]byte("TEST EVENT 2")))
	ch.Send(ctx, model.NewMessage([]byte("TEST EVENT 3")))

	time.Sleep(time.Millisecond * 50)
}
