package repository

import (
	"context"
	"dttg/internal/model"
)

type Repository interface {
	GetByName(ctx context.Context, ID string) (*model.Class, error)
	GetAll(ctx context.Context) ([]model.Class, error)
}
