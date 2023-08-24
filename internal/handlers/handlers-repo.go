package handlers

import (
	"github.com/mrkouhadi/chinauptodate/internal/config"
	db "github.com/mrkouhadi/chinauptodate/internal/db/postgres"
)

// repository type

type Repository struct {
	App *config.AppConfig
	DB  db.PgxRepository
}

// the repo that handlers methods belong to
var Repo *Repository

// NewRepo creates the new repository
func NewRepo(a *config.AppConfig, db *db.PgxRepository) *Repository {
	return &Repository{
		App: a,
		DB:  *db,
	}
}

// NewHandlers sets the repository for the handlers
func NewHandlers(r *Repository) {
	Repo = r
}
