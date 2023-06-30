package handlers

import "github.com/mrkouhadi/chinauptodate/internal/config"

// repository type

type Repository struct {
	App *config.AppConfig
}

// the repo that handlers methods belong to
var Repo *Repository

// NewRepo creates the new repository
func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
