package db

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mrkouhadi/chinauptodate/internal/db"
	"github.com/mrkouhadi/chinauptodate/internal/helpers"
	"github.com/mrkouhadi/chinauptodate/internal/models"
)

// ////////////////////////////////////////////Register a user
func (r *PgxRepository) InsertUser(a models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	statement := `INSERT INTO users(user_id, first_name,last_name, gender, email, password, profile_image, last_login, created_at, updated_at)
							  values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`

	// insert data into the database (Postgresql)
	hashedPassword, _ := helpers.HashPassword(a.Password)
	_, err := r.db.Exec(ctx, statement, a.User_id, a.First_name, a.Last_name, a.Gender, a.Email, hashedPassword, a.Profile_image, a.Last_login, a.Created_at, a.Updated_at)
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

// ////////////////// Authenticate Users
func (r *PgxRepository) AuthenticateUser(email, password string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	row := r.db.QueryRow(ctx, "SELECT * FROM users WHERE email = $1", email)

	var a models.User
	if err := row.Scan(&a.User_id, &a.First_name, &a.Last_name, &a.Gender, &a.Email, &a.Password, &a.Profile_image, &a.Last_login, &a.Created_at, &a.Updated_at); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, db.ErrNotExist
		}
		return nil, err
	}
	if !helpers.CheckPasswordHash(password, a.Password) {
		return nil, errors.New("incorrect credentials")
	}
	return &a, nil
}

// update user
func (r *PgxRepository) UpdateUser(id uuid.UUID, updated models.User) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	query := "UPDATE users SET first_name=$1,last_name=$2, gender=$3, email=$4, password=$5, profile_image=$6, last_login=$7, updated_at=$8 WHERE user_id = $9"
	res, err := r.db.Exec(ctx, query, updated.First_name, updated.Last_name, updated.Gender, updated.Email, updated.Password, updated.Profile_image, updated.Last_login, updated.Updated_at, id)
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
