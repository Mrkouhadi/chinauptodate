package handlers

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) ArticleGet(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	category := chi.URLParam(r, "category")
	_ = chi.URLParam(r, "id") // id of the article
	stringMap["title"] = category
	//
	id := uuid.MustParse("523673b4-8ca0-402f-a5fe-d71e8d485973")
	art, err := handlerRepo.DB.GetArticleByID(id)
	if err != nil {
		fmt.Println(err)
		handlerRepo.App.Session.Put(r.Context(), "error", "Could not Insert author into the database, Try again")
	}
	fmt.Println("article: ", art)
	// render html template
	render.Template(w, r, "article.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
