package main

import (
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

	// ROUTES
	r.Get("/test", Handler.TestUser)

	r.Get("/", Handler.Home)

	return r
}
