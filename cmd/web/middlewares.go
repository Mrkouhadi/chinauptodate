package main

import (
	"net/http"

	"github.com/justinas/nosurf"
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
