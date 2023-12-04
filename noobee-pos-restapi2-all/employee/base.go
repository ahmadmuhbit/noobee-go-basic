package employee

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Register(router *chi.Mux, db *sql.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	router.Route("/api/employees", func(r chi.Router) {
		r.Get("/", handler.ListEmployees)
		r.Get("/{id}", handler.GetById)
		r.Post("/", handler.Add)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})
}
