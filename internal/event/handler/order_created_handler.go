package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"github.com/jorgemarinho/go-expert-clean-architecture/pkg/events"
	"github.com/streadway/amqp"
)

type OrderCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *OrderCreatedHandler {
	return &OrderCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (h OrderCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Order created: %v", event.GetPayload())
	jsonOutput, err := json.Marshal(event.GetPayload())
	if err != nil {
		panic(err)
	}
	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	h.RabbitMQChannel.Publish(
		"amq.direct", //exchange
		"",           //key name
		false,        //mandatory
		false,        //immediate
		msgRabbitmq,  //message to publish
	)
}
