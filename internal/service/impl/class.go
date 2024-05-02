package impl

import (
	"context"
	"dttg/internal/model"
	"dttg/internal/repository"
)

type ClassService struct {
	repo repository.Repository
}

func NewClassService(repo repository.Repository) *ClassService {
	return &ClassService{repo: repo}
}

func (bs ClassService) GetByClassName(ctx context.Context, ID string) (*model.Class, error) {
	return bs.repo.GetByName(ctx, ID)
}

func (bs ClassService) GetAllClasses(ctx context.Context) ([]model.Class, error) {
	return bs.repo.GetAll(ctx)
}
