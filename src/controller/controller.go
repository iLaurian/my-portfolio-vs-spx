package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/iLaurian/my-portfolio-vs-spx/entity"
	"github.com/iLaurian/my-portfolio-vs-spx/service"
)

type Controller struct {
	TransactionService service.TransactionService
	HoldingService     service.HoldingService
}

type Config struct {
	R                  *gin.Engine
	TransactionService service.TransactionService
	HoldingService     service.HoldingService
}

func NewController(c *Config) {
	controller := &Controller{
		TransactionService: c.TransactionService,
		HoldingService:     c.HoldingService,
	}

	apiRoutes := c.R.Group("/api")
	{
		apiRoutes.GET("/txn/:id", controller.FindTransactionById)
		apiRoutes.GET("/txn", controller.FindAllTransactions)
		apiRoutes.POST("/txn/add", controller.AddTransaction)
		apiRoutes.POST("/txn/edit", controller.EditTransaction)
		apiRoutes.DELETE("/txn/delete", controller.DeleteTransaction)

		apiRoutes.GET("/hldg", controller.GetAllHoldings)
		apiRoutes.GET("/hldg/update", controller.UpdateAllHoldings)
		apiRoutes.DELETE("/hldg/delete", controller.DeleteAllHoldings)
	}
}

func (c *Controller) FindTransactionById(ctx *gin.Context) {
	transaction, err := c.TransactionService.FindById(ctx.Request.Context(), ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"transaction": transaction})
	}
}

func (c *Controller) FindAllTransactions(ctx *gin.Context) {
	transactions, err := c.TransactionService.FindAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{"transactions": transactions})
	}
}

func (c *Controller) AddTransaction(ctx *gin.Context) {
	var txn entity.Transaction
	log.Println(ctx)
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.TransactionService.Add(ctx.Request.Context(), txn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{"message": "Added successfully"})
}

func (c *Controller) EditTransaction(ctx *gin.Context) {
	var txn entity.Transaction
	if err := ctx.ShouldBindJSON(&txn); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.TransactionService.Edit(ctx.Request.Context(), txn); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Edited successfully"})
}

func (c *Controller) DeleteTransaction(ctx *gin.Context) {
	var request struct {
		ID int `json:"id"`
	}
	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.TransactionService.Delete(ctx.Request.Context(), request.ID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted successfully"})
}

func (c *Controller) GetAllHoldings(ctx *gin.Context) {
	holdings, err := c.HoldingService.GetAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusOK, gin.H{"holdings": holdings})
	}
}

func (c *Controller) UpdateAllHoldings(ctx *gin.Context) {
	holdings, err := c.HoldingService.UpdateAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"updated holdings": holdings})
	}
}

func (c *Controller) DeleteAllHoldings(ctx *gin.Context) {
	err := c.HoldingService.DeleteAll(ctx.Request.Context())
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted successfully"})
	}
}
