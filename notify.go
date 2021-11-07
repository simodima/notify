package notify

import (
	"context"
	"net/http"

	"github.com/go-redis/redis/v8"

	"github.com/toretto460/notify/internal/channel"
	"github.com/toretto460/notify/internal/driver"
)

// ChannelFactory is the factory used to create new channels instances
type ChannelFactory interface {
	New() (channel.Channel, error)
	Get(id string) (channel.Channel, error)
}

// Standalone creates new standalone channel factory.
// Beware that the standalone driver is
// meant to be used only for development purpose because cannot be
// scaled to multiple processes (such as containers).
func Standalone(ctx context.Context) *channel.Factory {
	driver := driver.NewStandalone(ctx)
	factory := channel.NewFactory(&driver)

	return &factory
}

// Redis creates a new redis channel factory.
// It requires a redis client to be used as backend for the driver.
func Redis(rCli *redis.Client) *channel.Factory {
	redisDriver := driver.NewRedis(rCli)
	factory := channel.NewFactory(&redisDriver)

	return &factory
}

// DefaultHandler is the default http handler to handle messages and
// send them to the client. The events can be named or anonymous
// respecting the Server sent events specification described at
// https://developer.mozilla.org/en-US/docs/Web/API/Server-sent_events
func DefaultHandler(factory *channel.Factory) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		qs := r.URL.Query()
		chID := qs.Get("channel")

		msgChan, err := factory.Get(chID)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		messages, err := msgChan.Receive(ctx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		for {
			select {
			case <-ctx.Done():
				return
			case m := <-messages:
				m.Write(w)
				// Flush the write to actually send data to the client
				if f, ok := w.(http.Flusher); ok {
					f.Flush()
				}
			}
		}
	}
}
