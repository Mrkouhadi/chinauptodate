package handlers

import (
	"github.com/mrkouhadi/chinauptodate/internal/config"
	"github.com/mrkouhadi/chinauptodate/internal/db"
)

// repository type

type Repository struct {
	App *config.AppConfig
	DB  db.PgxRepository // TODO:
}

// the repo that handlers methods belong to
var Repo *Repository

// NewRepo creates the new repository
func NewRepo(a *config.AppConfig, db *db.PgxRepository) *Repository {
	return &Repository{
		App: a,
		DB:  *db, // TODO:
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
