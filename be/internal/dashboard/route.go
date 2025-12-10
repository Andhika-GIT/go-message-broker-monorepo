package dashboard

import "github.com/go-chi/chi/v5"

func NewDashboardRoutes(r chi.Router, handler *DashboardHandler) {
	r.Get("/dashboard", handler.GetDataSummary)
}
