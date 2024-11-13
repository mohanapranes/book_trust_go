package server

import (
	"context"
	"fmt"
	"github/mohanapranes/book_trust_go/config"
	"github/mohanapranes/book_trust_go/config/database"
	"github/mohanapranes/book_trust_go/pkg/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func Start(conf *config.Config, database database.Database) {
	router := gin.Default()

	routes.RegisterRoutes(router, database)

	serverDetails := conf.Server
	adderss := fmt.Sprintf("%s:%d", serverDetails.Host, serverDetails.Port)

	server := http.Server{
		Addr:    adderss,
		Handler: router,
	}

	log.Printf("Starting server on %s", adderss)
	// Start the server in a goroutine
	go func() {
		log.Printf("Starting server on %s", adderss)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutting down server...")

	// Gracefully shut down with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited gracefully")
}
