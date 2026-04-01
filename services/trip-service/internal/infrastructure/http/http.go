package http

import (
	"encoding/json"
	"net/http"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/shared/types"
)

type HttpHandler struct {
	Service domain.TripService
}

type previewTripRequest struct {
	UserID      string           `json:"userID"`
	Pickup      types.Coordinate `json:"pickup"`
	Destination types.Coordinate `json:"destination"`
}

func (h *HttpHandler) HandleTripPreview(w http.ResponseWriter, r *http.Request) {
	var reqBody previewTripRequest
	if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
		http.Error(w, "invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	t, err := h.Service.CreateTrip(r.Context(), &domain.RideFareModel{
		UserID: "42",
	})
	if err != nil {
		http.Error(w, "failed to create trip", http.StatusInternalServerError)
		return
	}

	writeJSON(w, http.StatusOK, t)
}

func writeJSON(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}
