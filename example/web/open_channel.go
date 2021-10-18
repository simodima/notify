package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/toretto460/notify/channel"
)

type channelUseCases interface {
	New() (channel.Channel, error)
}

func NewOpenChannel(chUseCases channelUseCases) OpenChannel {
	return OpenChannel{
		useCase: chUseCases,
	}
}

type OpenChannel struct {
	useCase channelUseCases
}

func (e *OpenChannel) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ch, err := e.useCase.New()

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	events, err := ch.Receive(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "data: %s\n\n", ch.ID())
	flush(w)

	for {
		select {
		case ev := <-events:
			log.Printf("sending message to the client %+v", ev)
			ev.Write(w)
			flush(w)
		case <-r.Context().Done():
			return
		}
	}
}

func flush(w http.ResponseWriter) {
	if f, ok := w.(http.Flusher); ok {
		f.Flush()
	}
}
