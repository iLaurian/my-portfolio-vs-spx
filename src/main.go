package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iLaurian/my-portfolio-vs-spx/controller"
	"github.com/iLaurian/my-portfolio-vs-spx/service"
)

var (
	transactionService    service.TransactionService       = service.New()
	transactionController controller.TransactionController = controller.New(transactionService)
)

func main() {
	log.Println("Starting server ...")

	router := gin.Default()

	apiRoutes := router.Group("/api")
	{
		apiRoutes.GET("/txn", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, transactionController.FindAll())
		})

		apiRoutes.POST("/txn/add", func(ctx *gin.Context) {
			err := transactionController.Add(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusCreated, gin.H{"message": "Added successfully"})
			}
		})

		apiRoutes.POST("/txn/edit", func(ctx *gin.Context) {
			err := transactionController.Edit(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusAccepted, gin.H{"message": "Edited successfully"})
			}
		})

		apiRoutes.POST("/txn/delete", func(ctx *gin.Context) {
			err := transactionController.Delete(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusAccepted, gin.H{"message": "Deleted successfully"})
			}
		})
	}

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Graceful server shutdown

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
