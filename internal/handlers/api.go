package handlers

import (
	"github.com/go-chi/chi"
	chimiddle "github.com/go-chi/chi/middleware"
	"goapi/internal/middleware"
)

func Handler(r *chi.Mux) {
	//gloable middleware
	r.Use(chimiddle.StripSlashes)

	r.Route("/account", func(router chi.Router) {
		router.Use(middleware.Authoization)

		router.Get("/coins", GetCoinBalance)
	})
}
