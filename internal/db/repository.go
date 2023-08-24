package db

import (
	"errors"

	"github.com/google/uuid"
	"github.com/mrkouhadi/chinauptodate/internal/models"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type Repository interface {
	CreateArticle(a models.Article) (*models.Article, error)
	CreateCategory(c models.Category) (*models.Category, error)
	GetArticleByID(id uuid.UUID) (*models.Article, error)
}
