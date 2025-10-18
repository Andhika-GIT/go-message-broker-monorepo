package order

import "github.com/go-chi/chi/v5"

func NewOrderRoutes(r chi.Router, handler *OrderHandler) {
	r.Post("/order/upload", handler.UploadOrder)
}
