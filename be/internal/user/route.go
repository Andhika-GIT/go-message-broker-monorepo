package user

import "github.com/go-chi/chi/v5"

func NewUserRoutes(r chi.Router, handler *UserHandler) {
	r.Get("/user", handler.GetAllUsers)
	r.Post("/user/upload", handler.UploadUser)
}
