package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) ArticleGet(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	stringMap := make(map[string]string)
	stringMap["title"] = category
	id := chi.URLParam(r, "id") // id of the article
	id_uuid, _ := uuid.Parse(id)

	// get all articles
	articles, _ := handlerRepo.DB.AllArticles()
	fmt.Println("Original title of the 1st article: ", articles[0].News_title)
	//  get a single article
	art, _ := handlerRepo.DB.GetArticleByID(id_uuid)
	/// update article
	newArticle := models.Article{
		News_title:     "UPDATED the new title of this article",
		News_image:     art.News_image,
		News_content:   art.News_content,
		Comment_status: true,
		Updated_at:     time.Now(),
	}
	NewArticle, err := handlerRepo.DB.UpdateArticle(id_uuid, newArticle)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("The new pdated article title:", NewArticle.News_title)
	// render html template
	render.Template(w, r, "article.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
	})
}
