package shared

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
	"github.com/spf13/viper"
)

type RabbitMqProducer struct {
	Conn    *amqp091.Connection
	Channel *amqp091.Channel
}

func NewRabbitMqProducer(viper *viper.Viper) (*RabbitMqProducer, error) {
	conn, err := amqp091.Dial(viper.GetString("RABBITMQ_CONNECTION_URL"))

	if err != nil {
		panic(fmt.Errorf("fatal error connecting to rabbitmq: %w", err))

	}

	ch, err := conn.Channel()

	if err != nil {
		panic(fmt.Errorf("fatal error connecting to rabbitmq channel: %w", err))
	}

	return &RabbitMqProducer{
		Conn:    conn,
		Channel: ch,
	}, nil
}

func (c *RabbitMqProducer) Close() {
	if c.Channel != nil {
		_ = c.Channel.Close()
	}

	if c.Conn != nil {
		_ = c.Conn.Close()
	}
}

func (c *RabbitMqProducer) QueueDeclare() (amqp091.Queue, error) {
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

func (c *RabbitMqProducer) Publish() error {
	err := c.Channel.Publish(
		"go-exchange", "test", false, false,
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte("ini dari handler controller"),
		},
	)

	if err != nil {
		return fmt.Errorf("failed to publish message: %v", err)
	}

	return nil
}
