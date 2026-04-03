package main

import (
	"log"
	"net/http"
	"ride-sharing/shared/contracts"
	"ride-sharing/shared/util"

	"github.com/gorilla/websocket"
)

// We need to create a WebSocket upgrader to upgrade HTTP connections to WebSocket connections
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for simplicity, consider restricting in production
	},
}

// handleRidersWebSocket handles WebSocket connections for riders
func handleRidersWebSocket(w http.ResponseWriter, r *http.Request) {
	// First we need to upgrade the HTTP connection to a WebSocket connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		log.Printf("Missing userID in query parameters")
		return
	}

	// Here, we are reading messages from the WebSocket connection in a loop
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Printf("Received message from rider %s: %s", userID, string(message))
	}
}

// handleDriversWebSocket handles WebSocket connections for drivers
func handleDriversWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Failed to upgrade to WebSocket: %v", err)
		return
	}
	defer conn.Close()

	userID := r.URL.Query().Get("userID")
	if userID == "" {
		log.Printf("Missing userID in query parameters")
		return
	}

	packageSlug := r.URL.Query().Get("packageSlug")
	if packageSlug == "" {
		log.Printf("Missing packageSlug in query parameters")
		return
	}

	type Driver struct {
		Id             string `json:"id"`
		Name           string `json:"name"`
		ProfilePicture string `json:"profilePicture"`
		CarPlate       string `json:"carPlate"`
		PackageSlug    string `json:"packageSlug"`
	}

	msg := contracts.WSMessage{
		Type: "driver.cmd.register",
		Data: Driver{
			Id:             userID,
			Name:           "John Doe",
			ProfilePicture: util.GetRandomAvatar(1),
			CarPlate:       "ABC-123",
			PackageSlug:    packageSlug,
		},
	}

	// Here, we are sending a message via the WebSocket connection to the driver after they connect
	if err := conn.WriteJSON(msg); err != nil {
		log.Printf("Error writing message: %v", err)
		return
	}

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %v", err)
			break
		}
		log.Printf("Received message from rider %s: %s", userID, string(message))
	}
}
