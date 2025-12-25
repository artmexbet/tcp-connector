package main

import (
	"context"
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"tcp-conntector/internal/api"
	"tcp-conntector/internal/checker"
	"time"
)

func main() {
	// Configuration
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Dependencies
	checkService := checker.NewTCPChecker()
	handler := api.NewHandler(checkService)

	// Router setup
	mux := http.NewServeMux()
	mux.HandleFunc("/check", handler.HandleCheck)

	// Server setup
	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Graceful shutdown setup
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Starting server on :%s", port)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Server failed: %v", err)
		}
	}()

	// Wait for signal
	<-done
	log.Print("Server stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
