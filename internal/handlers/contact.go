package handlers

import (
	"net/http"

	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) ContactGet(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "Contact us"
	render.Template(w, r, "contact.page.gohtml", &models.TemplateData{StringMap: strMap})
}
