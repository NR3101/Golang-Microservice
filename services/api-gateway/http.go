package main

import (
	"encoding/json"
	"log"
	"net/http"
	"ride-sharing/services/api-gateway/grpc_clients"
	"ride-sharing/shared/contracts"
)

func handleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if reqBody.UserID == "" {
		log.Printf("missing userID in request body")
		http.Error(w, "missing userID in request body", http.StatusBadRequest)
		return
	}

	// Why to create a new gRPC client for each request: bcz if a service is down we don't want to block the entire API gateway, we can just return an error for that specific request and let the client handle it, instead of blocking all requests until the service is back up. This way we can have better fault tolerance and resilience in our system.
	tripService, err := grpc_clients.NewTripServiceClient()
	if err != nil {
		log.Printf("failed to create trip service client: %v", err)
		return
	}
	defer tripService.Close()

	tripPreview, err := tripService.Client.PreviewTrip(r.Context(), reqBody.ToProto())
	if err != nil {
		log.Printf("failed to get trip preview: %v", err)
		http.Error(w, "failed to get trip preview", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: tripPreview}

	writeJSON(w, http.StatusCreated, response)
}
