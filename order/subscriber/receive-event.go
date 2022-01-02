package subscriber

import "github.com/nats-io/nats.go"

type DataSub struct {
	EventName string
	Callback  func(status bool) int
}

func Subs(dataSub []DataSub) {
	nc, _ := nats.Connect("nats://nats:4222")

	for _, sub := range dataSub {
		nc.Subscribe(sub.EventName, func(m *nats.Msg) {
			sub.Callback(string(m.Data) == "true")
		})
	}
}
