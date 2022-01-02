package main

import (
	"io/ioutil"
	"log"
	"math/rand"

	"github.com/gofiber/fiber/v2"
	"github.com/nugrohosam/go-microservice-sagas/customer/publisher"
	"github.com/nugrohosam/go-microservice-sagas/customer/subscriber"
	"gopkg.in/yaml.v2"
)

type event struct {
	Order []string `yaml:"order"`
}

func (c *event) getEvent() *event {

	yamlFile, err := ioutil.ReadFile("events.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func doCallbackFunc() {
	currentRand := rand.Intn(10)
	isTrue := (currentRand % 2) == 0
	if isTrue {
		publisher.Pub("order-create", "true")
	} else {
		publisher.Pub("order-create", "false")
	}
}

func main() {
	app := fiber.New()

	currentEvent := event{}
	orderEvents := currentEvent.getEvent().Order
	dataEvents := make([]subscriber.DataSub, len(orderEvents))

	for index, data := range orderEvents {
		dataEvents[index] = subscriber.DataSub{data, doCallbackFunc}
	}

	subscriber.Subs(dataEvents)

	app.Listen(":3000")
}
