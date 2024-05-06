package impl

import (
	"context"
	"dttg/internal/repository"
)

type ClassService struct {
	repo repository.Repository
}

func NewClassService(repo repository.Repository) *ClassService {
	return &ClassService{repo: repo}
}

func (bs ClassService) GetByNameServ(ctx context.Context, Name string) (*interface{}, error) {
	return bs.repo.GetByName(ctx, Name)
}

func (bs ClassService) GetAllServ(ctx context.Context) ([]interface{}, error) {
	return bs.repo.GetAll(ctx)
}
