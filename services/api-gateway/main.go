package main

import (
	"fmt"
	"log"
	"net/http"

	"ride-sharing/shared/env"
)

var (
	httpAddr = env.GetString("HTTP_ADDR", ":8081")
)

func main() {
	log.Println("Starting API Gateway")

	mux := http.NewServeMux()

	mux.HandleFunc("POST /trip/preview", handleTripPreview)

	server := &http.Server{
		Addr:    httpAddr,
		Handler: mux,
	}

	fmt.Printf("API Gateway is running on %s\n", httpAddr)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("Failed to start server: %v", err)
	}

}
