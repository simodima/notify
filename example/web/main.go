package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/toretto460/notify/channel"
	"github.com/toretto460/notify/client"
)

var redisCli *redis.Client

func init() {
	redisCli = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	if err := redisCli.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}
}

func main() {
	redisClient := client.NewRedis(redisCli)
	chFactory := channel.NewFactory(&redisClient)
	events := NewOpenChannel(&chFactory)
	messages := NewSendMessage(&chFactory)

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.Handle("/open", &events)
	http.Handle("/notify", &messages)

	log.Print("Starting web server at :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
