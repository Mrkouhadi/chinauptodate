package db

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PgxRepository struct {
	db *pgxpool.Pool
}

// NewPgxRepository creates a new postgresql repository
func NewPgxRepository(db *pgxpool.Pool) *PgxRepository {
	return &PgxRepository{
		db: db,
	}
}

////////////////////////////////////////////// create a car data

func (r *PgxRepository) CreateCar(c Car) (*Car, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var id int64
	err := r.db.QueryRow(ctx, "INSERT INTO cars(brand, model, color, price) values($1, $2, $3, $4) RETURNING id", c.Brand, c.Model, c.Color, c.Price).Scan(&id)
	if err != nil {
		var pgxError *pgconn.PgError
		if errors.As(err, &pgxError) {
			if pgxError.Code == "23505" {
				return nil, ErrDuplicate
			}
		}
		return nil, err
	}
	c.ID = id

	return &c, nil
}
