package publisher

import "github.com/nats-io/nats.go"

func Pub(event string) {
	nc, _ := nats.Connect("nats://nats:4222")
	nc.Publish(event, []byte("Hello From Order"))
}
