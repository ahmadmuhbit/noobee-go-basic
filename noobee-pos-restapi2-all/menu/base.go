package menu

import (
	"database/sql"

	"github.com/go-chi/chi/v5"
)

func Register(router *chi.Mux, db *sql.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	router.Route("/api/menus", func(r chi.Router) {
		r.Get("/", handler.GetAll)
		r.Get("/{id}", handler.GetById)
		r.Post("/", handler.Create)
		r.Put("/{id}", handler.Update)
		r.Delete("/{id}", handler.Delete)
	})
}
