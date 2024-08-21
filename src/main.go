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
	"github.com/iLaurian/my-portfolio-vs-spx/repository"
	"github.com/iLaurian/my-portfolio-vs-spx/service"
)

func main() {
	log.Println("Starting server ...")

	db, err := initDB()

	if err != nil {
		log.Fatalf("Unable to initialize database: %v\n", err)
	}

	transactionRepository := repository.NewTransactionRepository(db.DB)
	transactionService := service.NewTransactionService(transactionRepository)

	holdingRepository := repository.NewHoldingRepository(db.RedisClient)
	holdingService := service.NewHoldingService(holdingRepository)

	router := gin.Default()

	controller.NewController(&controller.Config{
		R:                  router,
		TransactionService: transactionService,
		HoldingService:     holdingService,
	})

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

	if err := db.close(); err != nil {
		log.Fatalf("A problem occurred gracefully shutting down the database connection: %v\n", err)
	}

	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	<-ctx.Done()
	log.Println("timeout of 5 seconds.")
	log.Println("Server exiting")
}
