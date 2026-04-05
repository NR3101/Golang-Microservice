package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", enableCORS(handleTripPreview))
	mux.HandleFunc("POST /trip/start", enableCORS(handleTripStart))
	mux.HandleFunc("/ws/drivers", handleDriversWebSocket)
	mux.HandleFunc("/ws/riders", handleRidersWebSocket)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	/******************************** Code for graceful shutdown ********************************/

	serverErrs := make(chan error, 1)

	go func() {
		log.Printf("API Gateway listening on %s", httpAddr)
		serverErrs <- server.ListenAndServe()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	select {
	case err := <-serverErrs:
		log.Printf("API Gateway error: %v", err)

	case sig := <-shutdown:
		log.Printf("Received signal: %v. Shutting down API Gateway...", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("Error during shutting down server gracefully: %v", err)
			if err := server.Close(); err != nil {
				log.Printf("Error during forcefully closing server: %v", err)
			}
		} else {
			log.Println("API Gateway shutdown gracefully")
		}
	}
}
