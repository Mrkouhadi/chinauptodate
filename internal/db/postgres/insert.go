package db

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mrkouhadi/chinauptodate/internal/db"
	"github.com/mrkouhadi/chinauptodate/internal/models"
)

// //////////////////////////////////////////// create an Article to the databse(table:news)
func (r *PgxRepository) CreateArticle(a models.Article) (*models.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	statement := `INSERT INTO news(news_id, category_id,author_id, news_title, news_image, news_content, comment_status, created_at, updated_at) 
							  values($1, $2, $3, $4,$5,$6,$7,$8,$9)`
	_, err := r.db.Exec(ctx, statement, a.News_id, a.Category_id, a.Author_id, a.News_title, a.News_image, a.News_content, a.Comment_status, a.Created_at, a.Updated_at)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, db.ErrDuplicate
			}
		}
		return nil, err
	}
	return &a, nil
}

// ////////////////////////////////////////////  CreateCategory creates a category in the database(table: categories)
func (r *PgxRepository) CreateCategory(c models.Category) (*models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	statement := `INSERT INTO categories(category_id, category_title, category_description, created_at, updated_at) values($1, $2, $3, $4, $5)`
	_, err := r.db.Exec(ctx, statement, c.Category_id, c.Category_title, c.Category_description, c.Created_at, c.Updated_at)

	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, db.ErrDuplicate
			}
		}
		return nil, err
	}

	return &c, nil
}

// ////////////////////////////////////////////  CreateAuthor creates an author in the database(table: authors)
func (r *PgxRepository) CreateAuthor(au models.Author) (*models.Author, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	statement := `INSERT INTO authors(author_id, first_name, last_name,gender,birth_date,email,password,profile_image,last_login, created_at, updated_at) values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	_, err := r.db.Exec(ctx, statement, au.Author_id, au.First_name, au.Last_name, au.Gender, au.Birth_date, au.Email, au.Password, au.Profile_image, au.Last_login, au.Created_at, au.Updated_at)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, db.ErrDuplicate
			}
		}
		return nil, err
	}
	return &au, nil
}
