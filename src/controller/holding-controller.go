package controller

import (
	"github.com/iLaurian/my-portfolio-vs-spx/entity"
	"github.com/iLaurian/my-portfolio-vs-spx/service"
)

type HoldingController interface {
	GetAll() ([]entity.Holding, error)
	UpdateAll() ([]entity.Holding, error)
	DeleteAll() error
}

type holdingController struct {
	service service.HoldingService
}

func NewHoldingController(service service.HoldingService) HoldingController {
	return &holdingController{
		service: service,
	}
}

func (c holdingController) GetAll() ([]entity.Holding, error) {
	return c.service.GetAll()
}

func (c holdingController) UpdateAll() ([]entity.Holding, error) {
	return c.service.UpdateAll()
}

func (c holdingController) DeleteAll() error {
	return c.service.DeleteAll()
}
