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
