package publisher

import "github.com/nats-io/nats.go"

func Pub(event string, message string, urlNats string) {
	nc, _ := nats.Connect(urlNats)
	nc.Publish(event, []byte(message))
}
