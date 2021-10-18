package main

import (
	"context"

	"github.com/toretto460/notify/channel"
	"github.com/toretto460/notify/client"
	"github.com/toretto460/notify/model"
)

func main() {
	client := client.NewStandalone(context.TODO())
	channeFactory := channel.NewFactory(&client)

	ch, _ := channeFactory.New()

	ch.Send(context.TODO(), model.NewEvent([]byte("TEST EVENT")))
}
