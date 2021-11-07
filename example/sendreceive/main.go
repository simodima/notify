package main

import (
	"context"
	"log"
	"time"

	"github.com/toretto460/notify"
	"github.com/toretto460/notify/model"
)

func main() {
	ctx := context.TODO()
	channeFactory := notify.Standalone(ctx)

	ch, _ := channeFactory.New()

	go func() {
		// Receive messages
		messages, _ := ch.Receive(ctx)
		for msg := range messages {
			log.Printf("Message received NAME: %s DATA: %s", msg.Name(), string(msg.Data()))
		}
	}()

	// Publish messages
	ch.Send(ctx, model.NewMessage("test", []byte("TEST EVENT 1")))
	ch.Send(ctx, model.NewMessage("test", []byte("TEST EVENT 2")))
	ch.Send(ctx, model.NewMessage("test", []byte("TEST EVENT 3")))

	time.Sleep(time.Millisecond * 50)
}
