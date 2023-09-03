package handlers

import (
	"net/http"

	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) ProfileGet(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	data := make(map[string]interface{})
	u := handlerRepo.App.Session.Get(r.Context(), "user")
	user := u.(models.User)
	data["user"] = user
	strMap["title"] = "Welcome" + user.First_name + "To your Profile"
	render.Template(w, r, "profile.page.gohtml", &models.TemplateData{
		StringMap: strMap,
		Data:      data,
	})
}
