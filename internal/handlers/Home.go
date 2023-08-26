package handlers

import (
	"net/http"

	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) Home(w http.ResponseWriter, r *http.Request) {
	strMap := make(map[string]string)
	strMap["title"] = "𝗦𝘁𝗮𝘆 𝘂𝗽𝗱𝗮𝘁𝗲𝗱 𝘄𝗶𝘁𝗵 𝗮𝗹𝗹 𝗢𝗻𝗴𝗼𝗶𝗻𝗴 𝗮𝗻𝗱 𝗨𝗽𝗰𝗼𝗺𝗶𝗻𝗴 𝗲𝘃𝗲𝗻𝘁𝘀 & 𝗻𝗲𝘄𝘀 𝗶𝗻 𝗠𝗮𝗶𝗻𝗹𝗮𝗻𝗱"

	arts, _ := handlerRepo.DB.AllArticles()
	cats, _ := handlerRepo.DB.AllCategories()

	type ArtsWithCat struct {
		Article       models.Article
		CategoryTitle string
	}
	allArtsWithCat := []ArtsWithCat{}
	for _, art := range arts {
		for _, cat := range cats {
			if cat.Category_id == art.Category_id {
				allArtsWithCat = append(allArtsWithCat, ArtsWithCat{art, cat.Category_title})
			}
		}
	}
	data := make(map[string]interface{})
	data["articles"] = allArtsWithCat

	render.Template(w, r, "home.page.gohtml", &models.TemplateData{
		StringMap: strMap,
		Data:      data,
	})
}
