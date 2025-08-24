package shared

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type RabbitMqConsumer struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

func NewRabbitMqConsumer(viper *viper.Viper) (*RabbitMqConsumer, error) {
	conn, err := amqp091.Dial(viper.GetString("RABBITMQ_CONNECTION_URL"))

	if err != nil {
		panic(fmt.Errorf("fatal error connecting to rabbitmq: %w", err))

	}

	ch, err := conn.Channel()

	if err != nil {
		panic(fmt.Errorf("fatal error connecting to rabbitmq channel: %w", err))
	}

	return &RabbitMqConsumer{
		Conn:    conn,
		Channel: ch,
	}, nil
}

func (c *RabbitMqConsumer) Close() {
	if c.Channel != nil {
		_ = c.Channel.Close()
	}

	if c.Conn != nil {
		_ = c.Conn.Close()
	}
}

func (c *RabbitMqConsumer) QueueDeclare() (amqp091.Queue, error) {
	q, err := c.Channel.QueueDeclare(
		"test_queue", true, false, false, false,
		amqp091.Table{
			"x-queue-type": "quorum",
		},
	)

	if err != nil {
		return q, fmt.Errorf("failed to decleare queue : %v", err.Error())
	}

	err = c.Channel.QueueBind(
		q.Name, "test", "go-exchange", false, nil,
	)

	if err != nil {
		return q, fmt.Errorf("failed to bind queue : %v", err.Error())
	}

	return q, nil
}

func (c *RabbitMqConsumer) Consume() (<-chan amqp091.Delivery, error) {

	q, err := c.QueueDeclare()

	if err != nil {
		return nil, err
	}

	msgs, err := c.Channel.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)

	if err != nil {
		return nil, fmt.Errorf("failed to consume message : %v", err.Error())
	}

	return msgs, nil
}
