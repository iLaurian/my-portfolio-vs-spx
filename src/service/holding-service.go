package service

import (
	"context"

	"github.com/iLaurian/my-portfolio-vs-spx/entity"
	"github.com/iLaurian/my-portfolio-vs-spx/repository"
)

type HoldingService interface {
	GetAll(ctx context.Context) ([]entity.Holding, error)
	UpdateAll(ctx context.Context) ([]entity.Holding, error)
	DeleteAll(ctx context.Context) error
}

type holdingService struct {
	repository repository.HoldingRepository
}

func NewHoldingService(repository repository.HoldingRepository) HoldingService {
	return &holdingService{
		repository: repository,
	}
}

func (s *holdingService) GetAll(ctx context.Context) ([]entity.Holding, error) {
	return nil, nil
}

func (s *holdingService) UpdateAll(ctx context.Context) ([]entity.Holding, error) {
	return nil, nil
}

func (s *holdingService) DeleteAll(ctx context.Context) error {
	return nil
}
