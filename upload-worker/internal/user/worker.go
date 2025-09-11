package user

import (
	"encoding/json"
	"log"

	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
)

func StartWorker(rmq *shared.RabbitMqConsumer) {
	defer rmq.Close()

	msgs, err := rmq.Consume(shared.QueueUserDirectImport)

	if err != nil {
		log.Println(err)
	}

	var users []UserImport
	for msg := range msgs {
		err := json.Unmarshal(msg.Body, &users)

		if err != nil {
			log.Println("error when converting", err.Error())
			continue
		}

		log.Printf("users is %v", users)

		for _, user := range users {
			log.Printf("Name: %s, Email: %s, Phone: %s\n", user.Name, user.Email, user.PhoneNumber)

		}
	}

}
