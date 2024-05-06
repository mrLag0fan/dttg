package impl

import (
	"context"
	"dttg/internal/repository"
)

type SpellService struct {
	repo repository.Repository
}

func NewSpellService(repo repository.Repository) *SpellService {
	return &SpellService{repo: repo}
}

func (bs SpellService) GetByNameServ(ctx context.Context, Name string) (*interface{}, error) {
	return bs.repo.GetByName(ctx, Name)
}

func (bs SpellService) GetAllServ(ctx context.Context) ([]interface{}, error) {
	return bs.repo.GetAll(ctx)
}
