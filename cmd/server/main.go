package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rafly-ananda/snappsy-uploader-api/internal/config"
	ginHttp "github.com/rafly-ananda/snappsy-uploader-api/internal/http"
	"github.com/rafly-ananda/snappsy-uploader-api/internal/storage"
)

func main() {
	cfg := config.Load()
	fmt.Print(cfg.MinioPresignedExpiry)

	minioStore, err := storage.NewMinio(cfg.MinIOEndpoint, cfg.MinIOAccessKey, cfg.MinIOSecretKey, cfg.MinIOBucket, cfg.MinioPresignedExpiry, false)

	if err != nil {
		log.Fatalf("failed to init minio: %v", err)
	}

	r := ginHttp.NewRouter(minioStore)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: r,
	}

	// Start server
	go func() {
		log.Printf("starting server on :%s", cfg.AppPort)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown error: %v", err)
	}
	log.Println("bye")
}
