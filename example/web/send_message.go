package main

import (
	"io/ioutil"
	"net/http"

	"github.com/toretto460/notify"
	"github.com/toretto460/notify/model"
)

func NewSendMessage(factory notify.ChannelFactory) SendMessage {
	return SendMessage{
		factory: factory,
	}
}

type SendMessage struct {
	factory notify.ChannelFactory
}

func (e *SendMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	chID := qs.Get("channel")

	ch, err := e.factory.Get(chID)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	payload, err := ioutil.ReadAll(r.Body)

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	err = ch.Send(r.Context(), model.NewMessage("test", payload))

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
