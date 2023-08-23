package handlers

import (
	"net/http"
	"time"

	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) ArticleGet(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	// category := chi.URLParam(r, "category")
	// id := chi.URLParam(r, "id")
	a := GetArticle(3)
	stringMap["title"] = a.News_title
	data := make(map[string]interface{})
	data["article_title"] = a.News_title
	data["article_image"] = a.News_image
	data["article"] = a

	// render html template
	render.Template(w, r, "article.page.gohtml", &models.TemplateData{StringMap: stringMap, Data: data})
}

type Article struct {
	News_id        int
	Category_id    int
	Author_id      int
	News_title     string
	News_image     string
	News_content   string
	Comment_status bool
	Created_at     time.Time
	Updated_at     time.Time
}

func GetArticle(id int) Article {
	for _, val := range articles {
		if val.News_id == id {
			return val
		}
	}
	return Article{}
}

var img = "https://i.guim.co.uk/img/media/a7fe7170defa865d2b96b829f05c5d8fa82d8edf/0_20_2201_1321/master/2201.jpg?width=1200&height=900&quality=85&auto=format&fit=crop&s=410f9319cd3f76b9dccd90c5b5bbad5f"
var articles = []Article{
	{1, 1, 1, "The best player in the world", img, "some content here", true, time.Now(), time.Now()},
	{2, 2, 2, "Financial potential opportunity", img, "some content here", false, time.Now(), time.Now()},
	{3, 3, 3, "Multi potential chance", img, "some content here", false, time.Now(), time.Now()},
}
