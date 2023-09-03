package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

// GET
func (handlerRepo *Repository) LoginGET(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "Login in to your account"
	render.Template(w, r, "login.page.gohtml", &models.TemplateData{StringMap: strMap})
}

// POST
func (handlerRepo *Repository) LoginPost(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "Login in to your account"
	password := r.Form.Get("password")
	email := r.Form.Get("email")

	u, err := handlerRepo.DB.AuthenticateUser(email, password)

	// wrong credentials:
	if err != nil {
		handlerRepo.App.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/user/login", http.StatusSeeOther)
		return
	}

	// correct credentials:
	uu := models.User{
		User_id:       u.User_id,
		First_name:    u.First_name,
		Last_name:     u.Last_name,
		Gender:        u.Gender,
		Email:         u.Email,
		Password:      u.Password,
		Last_login:    time.Now(),
		Profile_image: u.Profile_image,
		Created_at:    u.Created_at,
		Updated_at:    u.Updated_at,
	}
	newU, err := handlerRepo.DB.UpdateUser(uu.User_id, uu)
	if err != nil {
		log.Println("could not update user last login", err)
		return
	}
	handlerRepo.App.Session.Put(r.Context(), "user", newU)
	http.Redirect(w, r, "/user/profile", http.StatusSeeOther)
}
