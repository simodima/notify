package main

import (
	"io/ioutil"
	"net/http"

	"github.com/toretto460/notify/channel"
	"github.com/toretto460/notify/model"
)

type channelOpener interface {
	Get(id string) (channel.Channel, error)
}

func NewSendMessage(chUseCase channelOpener) SendMessage {
	return SendMessage{
		useCase: chUseCase,
	}
}

type SendMessage struct {
	useCase channelOpener
}

func (e *SendMessage) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()
	chID := qs.Get("channel")

	ch, err := e.useCase.Get(chID)

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

	err = ch.Send(r.Context(), model.NewMessage(payload))

	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
