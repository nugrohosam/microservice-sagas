package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/nugrohosam/go-microservice-sagas/order/publisher"
	"github.com/nugrohosam/go-microservice-sagas/order/subscriber"
	"gopkg.in/yaml.v2"
)

type event struct {
	Customer []string `yaml:"customer"`
}

func (c *event) getEvent() *event {

	yamlFile, err := ioutil.ReadFile("./events.yaml")
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	err = yaml.Unmarshal(yamlFile, c)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return c
}

func doCallbackFunc(isComplete bool) int {
	if isComplete {
		fmt.Println("Complete jokes")
	} else {
		fmt.Println("Non complete jokes")
	}

	return 0
}

func main() {
	app := fiber.New()

	currentEvent := event{}
	customerEvents := currentEvent.getEvent().Customer
	customerEventsLength := len(customerEvents)
	dataEvents := make([]subscriber.DataSub, customerEventsLength)

	for index, item := range customerEvents {
		dataEvents[index] = subscriber.DataSub{item, doCallbackFunc}
	}

	subscriber.Subs(dataEvents)

	app.Post("/", func(c *fiber.Ctx) error {
		publisher.Pub("customer-event")
		return c.SendString("Process order")
	})

	app.Listen(":3000")
}
