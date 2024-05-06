package service

import (
	"context"
)

type Service interface {
	GetByNameServ(ctx context.Context, Name string) (*interface{}, error)
	GetAllServ(ctx context.Context) ([]interface{}, error)
}
