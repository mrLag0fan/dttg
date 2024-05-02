package service

import (
	"context"
	"dttg/internal/model"
)

type Service interface {
	GetByClassName(ctx context.Context, ID string) (*model.Class, error)
	GetAllClasses(ctx context.Context) ([]model.Class, error)
}
