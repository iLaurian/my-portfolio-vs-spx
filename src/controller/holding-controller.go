package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iLaurian/my-portfolio-vs-spx/entity"
	"github.com/iLaurian/my-portfolio-vs-spx/service"
)

type HoldingController interface {
	GetAll(ctx *gin.Context) ([]entity.Holding, error)
	UpdateAll(ctx *gin.Context) ([]entity.Holding, error)
	DeleteAll(ctx *gin.Context) error
}

type holdingController struct {
	service service.HoldingService
}

func NewHoldingController(service service.HoldingService) HoldingController {
	return &holdingController{
		service: service,
	}
}

func (c holdingController) GetAll(ctx *gin.Context) ([]entity.Holding, error) {
	return c.service.GetAll(ctx.Request.Context())
}

func (c holdingController) UpdateAll(ctx *gin.Context) ([]entity.Holding, error) {
	return c.service.UpdateAll(ctx.Request.Context())
}

func (c holdingController) DeleteAll(ctx *gin.Context) error {
	return c.service.DeleteAll(ctx.Request.Context())
}
