package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	h "ride-sharing/services/trip-service/internal/infrastructure/http"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"ride-sharing/shared/env"
	"syscall"
	"time"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8083")
)

func main() {
	inmemRepo := repository.NewInMemRepository()
	svc := service.NewService(inmemRepo)

	mux := http.NewServeMux()

	httpHandler := &h.HttpHandler{
		Service: svc,
	}

	mux.HandleFunc("POST /preview", httpHandler.HandleTripPreview)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	serverErr := make(chan error, 1)

	go func() {
		log.Printf("Trip Service listening on %s", httpAddr)
		serverErr <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErr:
		log.Printf("Trip Service error: %v", err)

	case sig := <-shutdown:
		log.Printf("Received signal: %v. Shutting down Trip Service...", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error during shutting down server gracefully: %v", err)
			if err := server.Close(); err != nil {
				log.Printf("Error during forcefully closing server: %v", err)
			}
		} else {
			log.Println("Trip Service shutdown gracefully")
		}
	}
}
