package handlers

import (
	"net/http"

	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) Article(w http.ResponseWriter, r *http.Request) {
	// stringMap := make(map[string]string)
	// stringMap["test"] = "test"
	// render.Template(w, r, "article.page.gohtml", &models.TemplateData{StringMap: stringMap})
	render.Template(w, r, "article.page.gohtml", &models.TemplateData{})

}
