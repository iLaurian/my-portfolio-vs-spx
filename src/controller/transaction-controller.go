package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iLaurian/my-portfolio-vs-spx/entity"
	"github.com/iLaurian/my-portfolio-vs-spx/service"
)

type TransactionController interface {
	FindAll(ctx *gin.Context) ([]entity.Transaction, error)
	Add(ctx *gin.Context) error
	Edit(ctx *gin.Context) error
	Delete(ctx *gin.Context) error
}

type transactionController struct {
	service service.TransactionService
}

func NewTransactionController(service service.TransactionService) TransactionController {
	return &transactionController{
		service: service,
	}
}

func (c transactionController) FindAll(ctx *gin.Context) ([]entity.Transaction, error) {
	return c.service.FindAll(ctx.Request.Context())
}

func (c transactionController) Add(ctx *gin.Context) error {
	var txn entity.Transaction
	err := ctx.ShouldBindJSON(&txn)
	if err != nil {
		return err
	}
	c.service.Add(ctx.Request.Context(), txn)
	return nil
}

func (c transactionController) Edit(ctx *gin.Context) error {
	var txn entity.Transaction
	err := ctx.ShouldBindJSON(&txn)
	if err != nil {
		return err
	}
	c.service.Edit(ctx.Request.Context(), txn)
	return nil
}

func (c transactionController) Delete(ctx *gin.Context) error {
	var request struct {
		ID int `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		return err
	}
	c.service.Delete(ctx.Request.Context(), request.ID)
	return nil
}
