package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/mrkouhadi/chinauptodate/internal/config"
	"github.com/mrkouhadi/chinauptodate/internal/handlers"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	// middlewares
	mux.Use(middleware.Recoverer)

	// GET routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/article", handlers.Repo.Article)

	// POST handlers

	// render STATIC FILES
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// return our mux
	return mux
}
