package handlers

import (
	"net/http"

	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) About(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "About us"
	render.Template(w, r, "about.page.gohtml", &models.TemplateData{StringMap: strMap})
}
