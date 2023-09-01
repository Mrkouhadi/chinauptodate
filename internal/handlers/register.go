package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/mrkouhadi/chinauptodate/internal/forms"
	"github.com/mrkouhadi/chinauptodate/internal/helpers"
	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

// GET
func (handlerRepo *Repository) RegisterGET(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "Register On our platform"
	render.Template(w, r, "register.page.gohtml", &models.TemplateData{
		StringMap: strMap,
		Form:      forms.New(nil),
	})
}

// POST
func (handlerRepo *Repository) RegisterPost(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println("ParseForm error: ", err)
		return
	}
	// get data from from
	fname := r.Form.Get("fname")
	lname := r.Form.Get("lname")
	email := r.Form.Get("email")
	password := r.Form.Get("password")
	gender := r.Form.Get("gender")
	// upload an image to the filesystem (./static/images)
	img_details, _ := helpers.UploadFile(w, r, email)

	// Submit a title of the page
	strMap := make(map[string]string)
	strMap["title"] = "Register On our platform"
	// set up the form errors and validation
	form := forms.New(r.PostForm)
	form.Has("fname", r)
	form.Has("lname", r)
	form.Has("gender", r)
	form.Has("email", r)
	form.Has("password", r)
	form.Has("confirm-password", r)
	form.Required("fname", "lname", "gender", "email", "password", "confirm-password")
	form.MinLength("password", 8)
	form.MinLength("fname", 4)
	form.MinLength("lname", 4)
	form.IsPassswordMatched("password", "confirm-password")

	// render the template
	if !form.Valid() { // when there errors
		render.Template(w, r, "register.page.gohtml", &models.TemplateData{
			StringMap: strMap,
			Form:      form,
		})
		return
	}

	newUser := models.User{
		User_id:       uuid.New(),
		First_name:    fname,
		Last_name:     lname,
		Email:         email,
		Password:      password,
		Gender:        gender,
		Profile_image: email + img_details.Extension,
		Last_login:    time.Now(),
		Created_at:    time.Now(),
		Updated_at:    time.Now(),
	}
	u, err := handlerRepo.DB.InsertUser(newUser)
	if err != nil {
		log.Println("error persisting data in the DB::", err)
		handlerRepo.App.Session.Put(r.Context(), "error", "Error registering, please try again later!")
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	log.Println(u)
	// when everything is okay
	handlerRepo.App.Session.Put(r.Context(), "flash", "You have been registered successfully")
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}
