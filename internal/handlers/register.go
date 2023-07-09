package handlers

import (
	"net/http"

	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) RegisterGET(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "register.page.gohtml", &models.TemplateData{})
}
