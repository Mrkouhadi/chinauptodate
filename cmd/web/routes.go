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
	mux.Use(LoadSession)
	mux.Use(NoSurf)
	//authentication Routes
	mux.Get("/register", handlers.Repo.RegisterGET)
	mux.Get("/login", handlers.Repo.LoginGET)
	// GET routes
	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)
	mux.Get("/contact", handlers.Repo.ContactGet)
	mux.Get("/{category}", handlers.Repo.CategoryGet)
	mux.Get("/{category}/{id}", handlers.Repo.ArticleGet)

	// admin
	mux.Get("/new-article", handlers.Repo.GETCreateNewArticle)
	mux.Post("/new-article", handlers.Repo.POSTnewArticle)

	// POST handlers PostSubscribe
	mux.Post("/", handlers.Repo.PostSubscribe)
	mux.Post("/register", handlers.Repo.RegisterPost)
	mux.Post("/login", handlers.Repo.LoginPost)

	// render STATIC FILES
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/*", http.StripPrefix("/static", fileServer))

	// return our mux
	return mux
}
