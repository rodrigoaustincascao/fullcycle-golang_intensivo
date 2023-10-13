package rabbitmq

import amqp "github.com/rabbitmq/amqp091-go"

func OpenChannel() (*amqp.Channel, error) {
	// Abre a conexão com rabbitmq
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		panic(err)
	}

	// Abre o canal
	ch, err := conn.Channel()
	// Indica quantas mensagem podem ser lidas por vez
	ch.Qos(100, 0, false)
	if err != nil {
		panic(err)
	}
	return ch, nil
}

// primeiro parâmetro é o channel de conexão com rabbitmq
func Consumer(ch *amqp.Channel, out chan amqp.Delivery) error {
	msgs, err := ch.Consume(
		"orders",
		"go-consumer",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		return err
	}

	for msg := range msgs {
		out <- msg
	}
	return nil
}
