package main

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/google/uuid"
	amqp "github.com/rabbitmq/amqp091-go"
)

type Order struct {
	ID    string
	Price float64
}

func Notify(ch *amqp.Channel, order Order) error {
	body, err := json.Marshal(order)
	if err != nil {
		return err
	}
	err = ch.Publish(
		"amq.direct",
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	return err
}

func GenerateOrders() Order {
	return Order{
		ID:    uuid.New().String(),
		Price: rand.Float64() * 100,
	}
}

func main() {
	// Abre a conexão com rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Abre o canal
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	// Enviando para o rabbitmq
	for i := 0; i < 100; i++ {
		order := GenerateOrders()
		err := Notify(ch, order)
		if err != nil {
			panic(err)
		}
		fmt.Println(order)
	}

}
