package main

import (
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/nugrohosam/go-microservice-sagas/customer/publisher"
	"github.com/nugrohosam/go-microservice-sagas/customer/subscriber"
	"gopkg.in/yaml.v2"
)

type event struct {
	Order    []string `yaml:"order"`
	Customer []string `yaml:"customer"`
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
	urlNats := "nats://" + os.Getenv("HOST_NATS")
	currentRand := rand.Intn(10)
	isTrue := (currentRand % 2) == 0
	publisher.Pub("customer-check-limit", strconv.FormatBool(isTrue), urlNats)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := fiber.New()

	currentEvent := event{}
	orderEvents := currentEvent.getEvent().Order
	dataEvents := make([]subscriber.DataSub, len(orderEvents))

	for index, data := range orderEvents {
		dataEvents[index] = subscriber.DataSub{data, doCallbackFunc}
	}

	urlNats := "nats://" + os.Getenv("HOST_NATS")
	subscriber.Subs(dataEvents, urlNats)

	port := ":" + os.Getenv("PORT")
	app.Listen(port)
}
