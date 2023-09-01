package handlers

import (
	"net/http"

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
	if err != nil {
		// wrong credentials
		handlerRepo.App.Session.Put(r.Context(), "error", "Invalid Credentials")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	handlerRepo.App.Session.Put(r.Context(), "user_id", u.User_id)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
