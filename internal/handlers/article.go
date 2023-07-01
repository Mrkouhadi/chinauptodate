package handlers

import (
	"net/http"

	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) Article(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "article.page.tmpl", &models.TemplateData{})
}
