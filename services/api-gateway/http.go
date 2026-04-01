package main

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
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
		http.Error(w, "userID is required", http.StatusBadRequest)
		return
	}

	jsonBody, _ := json.Marshal(reqBody)
	reader := bytes.NewReader(jsonBody)

	// Call trip service
	res, err := http.Post("http://trip-service:8083/preview", "application/json", reader)
	if err != nil {
		log.Printf("failed to call trip service: %v", err)
		return
	}
	defer res.Body.Close()

	var respBody any
	if err := json.NewDecoder(res.Body).Decode(&respBody); err != nil {
		http.Error(w, "failed to decode trip service response", http.StatusInternalServerError)
		return
	}

	response := contracts.APIResponse{Data: respBody}

	writeJSON(w, http.StatusCreated, response)
}
