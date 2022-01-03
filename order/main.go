package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nugrohosam/go-microservice-sagas/order/publisher"
	"github.com/nugrohosam/go-microservice-sagas/order/subscriber"
	"gopkg.in/yaml.v2"
)

type event struct {
	Order    []string `yaml:"order"`
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
		fmt.Println("Limit avail")
	} else {
		fmt.Println("Limit not avail")
	}

	return 0
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	currentEvent := event{}
	customerEvents := currentEvent.getEvent().Customer
	customerEventsLength := len(customerEvents)
	dataEvents := make([]subscriber.DataSub, customerEventsLength)

	for index, item := range customerEvents {
		dataEvents[index] = subscriber.DataSub{item, doCallbackFunc}
	}

	urlNats := "nats://" + os.Getenv("HOST_NATS")
	subscriber.Subs(dataEvents, urlNats)

	app.Post("/", func(c *fiber.Ctx) error {
		publisher.Pub("order-create", urlNats)
		return c.SendString("Process order")
	})

	port := ":" + os.Getenv("PORT")
	app.Listen(port)
}
