package publisher

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

func Pub(event string, urlNats string) {
	nc, _ := nats.Connect(urlNats)
	fmt.Println("published event", event)
	nc.Publish(event, []byte("Hello From Order"))
}
