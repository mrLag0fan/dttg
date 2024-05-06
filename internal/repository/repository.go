package repository

import (
	"context"
)

type Repository interface {
	GetByName(ctx context.Context, Name string) (*interface{}, error)
	GetAll(ctx context.Context) ([]interface{}, error)
}
