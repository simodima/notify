package main

import (
	"context"
	"log"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/toretto460/notify"
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
	chFactory := notify.Redis(redisCli)
	messages := NewSendMessage(chFactory)

	handler := notify.DefaultHandler(chFactory)
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

	http.HandleFunc("/open", handler)
	http.Handle("/notify", &messages)

	log.Print("Starting web server at :3000")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
