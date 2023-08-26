package handlers

import (
	"fmt"
	"net/http"

	"github.com/mrkouhadi/chinauptodate/internal/forms"
	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (repo *Repository) PostSubscribe(w http.ResponseWriter, r *http.Request) {
	// err := r.ParseForm()
	// if err != nil {
	// 	log.Println(err)
	// }
	email := r.FormValue("email")
	form := forms.New(r.PostForm)
	form.IsEmail(email)

	if !form.Valid() {
		repo.App.Session.Put(r.Context(), "error", "Could not subscribe, Your email is not valid please try again!")
		render.Template(w, r, "home.page.gohtml", &models.TemplateData{
			Form: form,
		})
		return
	}
	fmt.Println("All's good, No Errors", len(form.Errors))
	render.Template(w, r, "home.page.gohtml", &models.TemplateData{
		Form: form,
	})
}

// FIXME: still errors subscribing...
