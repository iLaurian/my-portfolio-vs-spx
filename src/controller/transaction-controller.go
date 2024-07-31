package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iLaurian/my-portfolio-vs-spx/entity"
	"github.com/iLaurian/my-portfolio-vs-spx/service"
)

type TransactionController interface {
	FindAll() []entity.Transaction
	Add(ctx *gin.Context) entity.Transaction
}

type controller struct {
	service service.TransactionService
}

func New(service service.TransactionService) TransactionController {
	return controller{
		service: service,
	}
}

func (c controller) FindAll() []entity.Transaction {
	return c.service.FindAll()
}

func (c controller) Add(ctx *gin.Context) entity.Transaction {
	var txn entity.Transaction
	ctx.BindJSON(&txn)
	c.service.Add(txn)
	return txn
}
