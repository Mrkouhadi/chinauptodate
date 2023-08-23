package db

import (
	"errors"
)

var (
	ErrDuplicate    = errors.New("record already exists")
	ErrNotExist     = errors.New("row does not exist")
	ErrUpdateFailed = errors.New("update failed")
	ErrDeleteFailed = errors.New("delete failed")
)

type Repository interface {
	CreateCar(car Car) (*Car, error)
}

type Car struct {
	ID    int64
	Brand string
	Model string
	Color string
	Price float64
}
