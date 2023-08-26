package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/mrkouhadi/chinauptodate/internal/forms"
)

func (repo *Repository) PostSubscribe(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
	email := r.FormValue("email")
	form := forms.New(r.PostForm)
	form.IsEmail(email)
	wrong := form.Valid() // false if no errors

	if !wrong { // something wrong
		fmt.Println("There are Errors", email, len(form.Errors))
		repo.App.Session.Put(r.Context(), "error", "Could not subscribe, Your email is not valid please try again!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	repo.App.Session.Put(r.Context(), "flash", "You have susbcribed to our newsletter successfully, Thank you !")
	fmt.Println("All's good, No Errors", email, len(form.Errors))
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
