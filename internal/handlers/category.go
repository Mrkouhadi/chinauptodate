package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

// GET: display all articles of a specific category
func (repoHandler *Repository) CategoryGet(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	category := chi.URLParam(r, "category")
	strMap["title"] = category
	arts, _ := repoHandler.DB.AllArticles()
	cats, _ := repoHandler.DB.AllCategories()
	var catsArticles []models.Article
	for _, cat := range cats {
		if cat.Category_title == category {
			for _, art := range arts {
				if art.Category_id == cat.Category_id {
					catsArticles = append(catsArticles, art)
				}
			}
			break
		}
	}
	data := make(map[string]interface{})
	data["articles"] = catsArticles
	render.Template(w, r, "category.page.gohtml", &models.TemplateData{
		StringMap: strMap,
		Data:      data,
	})
}
