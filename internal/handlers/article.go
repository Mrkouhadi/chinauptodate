package handlers

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/mrkouhadi/chinauptodate/internal/helpers"
	"github.com/mrkouhadi/chinauptodate/internal/models"
	"github.com/mrkouhadi/chinauptodate/internal/render"
)

func (handlerRepo *Repository) ArticleGet(w http.ResponseWriter, r *http.Request) {
	category := chi.URLParam(r, "category")
	stringMap := make(map[string]string)
	stringMap["title"] = category
	id := chi.URLParam(r, "id")
	id_uuid, _ := uuid.Parse(id)
	art, _ := handlerRepo.DB.GetArticleByID(id_uuid)

	data := make(map[string]interface{})
	data["article"] = art

	//createig a new
	// auId, _ := uuid.Parse("868f3b65-cb0f-4487-96fd-e425af6d868b")
	// catid, _ := uuid.Parse("0d23f565-e1b2-483f-89cb-e3d2cc4bd084")
	// article := models.Article{
	// 	News_id:     uuid.New(),
	// 	Category_id: catid,
	// 	Author_id:   auId,
	// 	News_title:  "Some International News to read",
	// 	News_image:  "./static/images/test.jpeg",
	// 	News_content: `générateur de magazines en ligne sont toujours populaires, et leurs auteurs jouissent de la célébrité et du respect. C’est pourquoi, pour de nombreux rédacteurs indépendants, écrire des articles dans des magazines est souvent un objectif de carrière – car la rémunération peut être dix fois plus élevée par mot que la rédaction d’articles ou de textes pour le journal local.
	// 	La rédaction d’articles de magazines exige des compétences différentes de celles des articles de blog, des scénarios ou des publicités. De plus, en tant que rédacteur de magazine, plus que dans toute autre industrie, vous devez vous spécialiser pour réussir. Vous écrivez des articles sur l’histoire différemment, sur le sport différemment, et sur l’histoire du sport d’une manière encore différente.
	// 	Le talent d’écriture, l’amour de la recherche méticuleuse et la souplesse dans la création de textes sont des compétences essentielles que vous devez maîtriser. C’est pourquoi de nombreuses personnes désireuses de créer et de publier leur propre magazine doivent maîtriser ce style spécifique et apprendre à rédiger un article de magazine.`,
	// 	Comment_status: true,
	// 	Created_at:     time.Now(),
	// 	Updated_at:     time.Now(),
	// }
	// addedArt, _ := Repo.DB.CreateArticle(article)
	// fmt.Println(addedArt.News_title)

	// creating a new category
	// cat := models.Category{
	// 	Category_id:          uuid.New(),
	// 	Category_title:       "GLOBAL",
	// 	Category_description: "All international news with all categories",
	// }
	// newcat, _ := handlerRepo.DB.CreateCategory(cat)
	// fmt.Println(newcat)
	// render html template
	render.Template(w, r, "article.page.gohtml", &models.TemplateData{
		StringMap: stringMap,
		Data:      data,
	})
}

// /////////////// ADMIN
// / insert a new article : http://localhost:8080/new-article
// GET
func (handlerRepo *Repository) GETCreateNewArticle(w http.ResponseWriter, r *http.Request) {
	render.Template(w, r, "new-article.page.gohtml", &models.TemplateData{})
}

// POST
func (handlerRepo *Repository) POSTnewArticle(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	// upload an image to the filesystem (./static/images)
	helpers.UploadFile(w, r, "new-image")

	// helpers.UploadMultipleFiles(w, r)
	handlerRepo.App.Session.Put(r.Context(), "flash", "An article has been added succesfully !")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
