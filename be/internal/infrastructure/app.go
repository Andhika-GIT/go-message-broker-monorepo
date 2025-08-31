package infrastructure

import (
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/order"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/shared"
	"github.com/Andhika-GIT/go-message-broker-monorepo/internal/user"
	"github.com/go-chi/chi/v5"
)

func InitApp() *chi.Mux {
	r := NewRouter()

	v := NewViper()
	rmq, _ := shared.NewRabbitMqProducer(v)
	order.NewOrderModule(r, rmq)
	user.NewUserModule(r, rmq)

	return r
}
