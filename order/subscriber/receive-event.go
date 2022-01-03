package subscriber

import (
	"fmt"

	"github.com/nats-io/nats.go"
)

type DataSub struct {
	EventName string
	Callback  func(status bool) int
}

func Subs(dataSub []DataSub, urlNats string) {
	nc, _ := nats.Connect(urlNats)

	for _, sub := range dataSub {
		fmt.Println("subscribed event", sub.EventName)
		nc.Subscribe(sub.EventName, func(m *nats.Msg) {
			sub.Callback(string(m.Data) == "true")
		})
	}
}
