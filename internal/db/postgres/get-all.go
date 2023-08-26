package db

import (
	"context"
	"time"

	"github.com/mrkouhadi/chinauptodate/internal/models"
)

// ////////////////////////////////// get all articles
func (r *PgxRepository) AllArticles() ([]models.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := r.db.Query(ctx, "SELECT * FROM news")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var allArticles []models.Article

	for rows.Next() {
		var article models.Article
		if err := rows.Scan(&article.News_id, &article.Category_id, &article.Author_id, &article.News_title, &article.News_image, &article.News_content, &article.Comment_status, &article.Created_at, &article.Updated_at); err != nil {
			return nil, err
		}
		allArticles = append(allArticles, article)
	}
	return allArticles, nil
}

// ////////////////////////////////// get all categories
func (r *PgxRepository) AllCategories() ([]models.Category, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rows, err := r.db.Query(ctx, "SELECT * FROM categories")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var allCategories []models.Category

	for rows.Next() {
		var category models.Category
		if err := rows.Scan(&category.Category_id, &category.Category_title, &category.Category_description, &category.Created_at, &category.Updated_at); err != nil {
			return nil, err
		}
		allCategories = append(allCategories, category)
	}
	return allCategories, nil
}
