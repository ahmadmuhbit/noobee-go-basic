package auth

import (
	"database/sql"
	routerChi "nobee-pos/infra/router/chi"

	"github.com/go-chi/chi/v5"
)

func Register(router *chi.Mux, db *sql.DB) {
	repo := NewRepository(db)
	svc := NewService(repo)
	handler := NewHandler(svc)

	router.Route("/api/auth", func(r chi.Router) {
		r.Post("/signup", handler.Register)
		r.Post("/signin", handler.Login)

		r.Group(func(r chi.Router) {
			r.Use(routerChi.CheckToken)
			r.Get("/profile", handler.Profile)
		})
	})
}
