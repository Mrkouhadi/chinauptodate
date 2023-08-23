package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (repoHandler *Repository) CategoryGet(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	// category := r.URL.Query().Get("category")
	category := chi.URLParam(r, "category")
	strMap["title"] = category
	fmt.Println(category)
	render.Template(w, r, "category.page.gohtml", &models.TemplateData{StringMap: strMap})
}
