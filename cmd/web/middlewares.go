package main

import (
	"net/http"

	"github.com/justinas/nosurf"
	"github.com/mrkouhadi/chinauptodate/internal/helpers"
)

// LoadSession loads and save session on every request
func LoadSession(next http.Handler) http.Handler {
	return session.LoadAndSave(next)
}

// NoSurf add CSRF protection to Post requests
func NoSurf(next http.Handler) http.Handler {
	// create a csrf handler
	CSRFHandler := nosurf.New(next)
	// set te base
	CSRFHandler.SetBaseCookie(http.Cookie{
		HttpOnly: true,
		Path:     "/",
		Secure:   app.InProduction, // will be changed in production cuz in development we don't use HTTPS we only http
		SameSite: http.SameSiteLaxMode,
	})
	return CSRFHandler
}

// user authentication
func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !helpers.IsUserAuthenticated(r) {
			session.Put(r.Context(), "error", "Please Log in first !")
			http.Redirect(w, r, "/user/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}
