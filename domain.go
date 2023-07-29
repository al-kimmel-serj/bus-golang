package bus

import (
	"context"

	"google.golang.org/protobuf/proto"
)

type (
	// EventFilter is used for filtering inbound events by EventKey. Event will be passed if event key starts with EventFilter value.
	EventFilter string

	// EventKey is used for group and route events. Events with same event key value must be consumed by same consumer.
	EventKey string

	// EventName determines type of event content.
	EventName string

	// EventVersion determines version of event content.
	EventVersion int

	// PublisherEndpoint is publisher network address in format `IP_ADDRESS:PORT`.
	PublisherEndpoint string
)

// Event contains data for publishing.
type Event[Payload proto.Message] struct {
	EventKey     EventKey
	EventPayload Payload
}

// Publisher writes events to events transport or to Subscriber's directly.
type Publisher[Payload proto.Message] interface {
	// Publish serializes passed events to protobuf and publishes passed events in same order as it was passed to method call.
	Publish(ctx context.Context, events []Event[Payload]) error
	// Stop releases resources allocated for Publisher.
	Stop() error
}

// PublishersRegistry allows to register/unregister and watch publishers which can write to Subscriber's directly.
type PublishersRegistry interface {
	// Register adds record about publisher endpoint to registry and returns unregister function which removes registered publisher from registry.
	Register(eventName EventName, eventVersion EventVersion, host string, port int) (unregister func() error, err error)
	// Watch creates watcher for publishers endpoints list updates and returns function for stop created watcher.
	Watch(eventName EventName, eventVersion EventVersion, handler func([]PublisherEndpoint)) (stop func() error, err error)
}

// Subscriber provides functions for subscribe/unsubscribe on specific EventName and EventVersion. One instance of Subscriber can't subscribe/unsubscribe on more than one pair of EventName and EventVersion.
type Subscriber[Payload proto.Message] interface {
	// EventsChan returns channel for receiving events matched with EventFilter.
	EventsChan() <-chan Event[Payload]
	// Stop removes subscriptions and releases resources allocated for Subscriber.
	Stop() error
	// Subscribe creates subscription for EventFilter.
	Subscribe(eventFilter EventFilter) error
	// Unsubscribe removes subscription for EventFilter.
	Unsubscribe(eventFilter EventFilter) error
}
