package bus

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type (
	EventFilter       string
	EventKey          string
	EventName         string
	EventVersion      int
	PublisherEndpoint string
)

type Event[Payload proto.Message] struct {
	EventKey     EventKey
	EventPayload Payload
}

type Publisher[Payload proto.Message] interface {
	Publish(ctx context.Context, events []Event[Payload]) error
	Stop() error
}

type PublishersRegistry interface {
	Register(eventName EventName, eventVersion EventVersion, host string, port int) (unregister func() error, err error)
	Watch(eventName EventName, eventVersion EventVersion, handler func([]PublisherEndpoint)) (stop func() error, err error)
}

type Subscriber[Payload proto.Message] interface {
	EventsChan() <-chan Event[Payload]
	Stop() error
	Subscribe(eventFilter EventFilter) error
	Unsubscribe(eventFilter EventFilter) error
}
