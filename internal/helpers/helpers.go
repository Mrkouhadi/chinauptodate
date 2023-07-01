package helpers

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/mrkouhadi/chinauptodate/internal/config"
)

var app *config.AppConfig

// Newhelpers sets up appconfig for helpers

func Newhelpers(a *config.AppConfig) {
	app = a
}

// ClientError sends errors to the client when something goes wrong on the clientside

func ClientError(w http.ResponseWriter, status int) {
	app.InfoLog.Println("Client Error With Status Of ", status)
	http.Error(w, http.StatusText(status), status)
}

// ServerError sends errors to the client when something goes wrong on the serverside
func ServerError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.ErrorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}
