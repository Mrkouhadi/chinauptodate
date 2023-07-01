package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mrkouhadi/chinauptodate/internal/config"
	"github.com/mrkouhadi/chinauptodate/internal/handlers"
	"github.com/mrkouhadi/chinauptodate/internal/helpers"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager
var infoLog *log.Logger
var errorLog *log.Logger

func main() {
	// set the project into DEV mode
	app.InProduction = false

	// set up error handling loggers (custom loggers for info and errors)
	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime) //  \t means tab (bunch of spaces)
	app.InfoLog = infoLog
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	app.ErrorLog = errorLog

	// set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	app.Session = session
	// set up db

	// create templates cache
	templateCache, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Cannot Create Template Cache !")
		return
	}
	app.TemplateCache = templateCache

	// define wether our app is using cached templates or not
	app.UseCache = false

	// set up implementation of handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	// start render of templates
	render.NewRenderer(&app)
	// start helpers
	helpers.Newhelpers(&app)
	// run the server
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	log.Printf("Listening on Port: %s", portNumber)
	err = server.ListenAndServe()
	log.Fatal(err)
}
