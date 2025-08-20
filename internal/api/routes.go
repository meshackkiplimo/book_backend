package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Routes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Route("/books", func(r chi.Router) {

	})

	return r
}
