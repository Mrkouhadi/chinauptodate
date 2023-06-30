package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/mrkouhadi/chinauptodate/internal/config"
	"github.com/mrkouhadi/chinauptodate/internal/handlers"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

const portNumber = ":8080"

var app config.AppConfig
var session *scs.SessionManager

func main() {
	// set the project into DEV mode
	app.InProduction = false
	// set up session
	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction
	// db

	// create templates cache
	tc, _ := render.CreateTemplateCache()
	// if err != nil {
	// 	log.Fatal("Cannot Create Template Cache !")
	// }
	app.TemplateCache = tc

	// define wether our app is using cached templates or not
	app.UseCache = false

	// set up implementation of handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)
	// start render of templates
	render.NewRenderer(&app)

	// run the server
	server := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}
	err := server.ListenAndServe()
	log.Fatal(err)
}
