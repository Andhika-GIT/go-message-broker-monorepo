package infrastructure

import "github.com/go-chi/chi/v5"

func InitApp() *chi.Mux {
	r := chi.NewRouter()

	v := NewViper()
	NewRabbitMqProducer(v)

	return r
}
