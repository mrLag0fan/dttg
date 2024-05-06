package impl

import (
	"context"
	"dttg/internal/repository"
)

type RaceService struct {
	repo repository.Repository
}

func NewRaceService(repo repository.Repository) *RaceService {
	return &RaceService{repo: repo}
}

func (bs RaceService) GetByNameServ(ctx context.Context, Name string) (*interface{}, error) {
	return bs.repo.GetByName(ctx, Name)
}

func (bs RaceService) GetAllServ(ctx context.Context) ([]interface{}, error) {
	return bs.repo.GetAll(ctx)
}
