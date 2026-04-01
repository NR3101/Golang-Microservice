package main

import (
	"encoding/json"
	"net/http"
)

// writeJSON is a helper function to write JSON responses with a given status code.
func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
