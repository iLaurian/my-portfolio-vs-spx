package service

import (
	"github.com/iLaurian/my-portfolio-vs-spx/entity"
	"github.com/iLaurian/my-portfolio-vs-spx/repository"
)

type HoldingService interface {
	GetAll() ([]entity.Holding, error)
	UpdateAll() ([]entity.Holding, error)
	DeleteAll() error
}

type holdingService struct {
	repository repository.HoldingRepository
}

func NewHoldingService(repository repository.HoldingRepository) HoldingService {
	return &holdingService{
		repository: repository,
	}
}

func (s *holdingService) GetAll() ([]entity.Holding, error) {
	return nil, nil
}

func (s *holdingService) UpdateAll() ([]entity.Holding, error) {
	return nil, nil
}

func (s *holdingService) DeleteAll() error {
	return nil
}
