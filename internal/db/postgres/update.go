package db

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mrkouhadi/chinauptodate/internal/db"
	"github.com/mrkouhadi/chinauptodate/internal/models"
)

// update an article
func (r *PgxRepository) UpdateArticle(id uuid.UUID, updated models.Article) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := r.db.Exec(ctx, "UPDATE news SET news_title=$1, news_image=$2, news_content=$3, comment_status=$4, updated_at=$5 WHERE news_id = $6", updated.News_title, updated.News_image, updated.News_content, updated.Comment_status, updated.Updated_at, id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, db.ErrDuplicate
			}
		}
		return nil, err
	}
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, db.ErrUpdateFailed
	}
	return &updated, nil
}

// update a category
func (r *PgxRepository) UpdateCategory(id uuid.UUID, updated models.Category) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := r.db.Exec(ctx, "UPDATE categories SET category_title=$1, category_description=$2, updated_at=$3 WHERE category_id = $4", updated.Category_title, updated.Category_description, updated.Updated_at, id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, db.ErrDuplicate
			}
		}
		return nil, err
	}
	rowsAffected := res.RowsAffected()
	if rowsAffected == 0 {
		return nil, db.ErrUpdateFailed
	}
	return &updated, nil
}
