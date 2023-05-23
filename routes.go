package main

import (
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func MakeRoutes(session *scs.SessionManager) chi.Router {
	r := chi.NewRouter()

	r.Use(middleware.RealIP)
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(session.LoadAndSave)

	// API Routes
	r.Route("/api", func(r chi.Router) {

	})

	// ROUTES
	r.Get("/test", Handler.TestUser)
	r.Get("/", Handler.Home)

	// Public file server
	fileServer := http.FileServer(http.Dir("./public"))
	r.Handle("/public/*", http.StripPrefix("/public", fileServer))

	return r
}
