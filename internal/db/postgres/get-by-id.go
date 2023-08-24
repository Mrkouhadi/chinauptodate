package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/mrkouhadi/chinauptodate/internal/db"
	"github.com/mrkouhadi/chinauptodate/internal/models"
)

// /////////////////////////////////// get a single article
func (r *PgxRepository) GetArticleByID(id uuid.UUID) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := r.db.QueryRow(ctx, "SELECT * FROM news WHERE news_id = $1", id)

	var article models.Article

	if err := row.Scan(&article.News_id, &article.Category_id, &article.Author_id, &article.News_title, &article.News_image, &article.News_content, &article.Comment_status, &article.Created_at, &article.Updated_at); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, db.ErrNotExist
		}
		return nil, err
	}
	return &article, nil
}
