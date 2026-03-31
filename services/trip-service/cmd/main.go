package main

import (
	"context"
	"fmt"
	"ride-sharing/services/trip-service/internal/domain"
	"ride-sharing/services/trip-service/internal/infrastructure/repository"
	"ride-sharing/services/trip-service/internal/service"
	"time"
)

func main() {
	inmemRepo := repository.NewInMemRepository()

	svc := service.NewService(inmemRepo)

	ctx := context.Background()
	fare := &domain.RideFareModel{
		UserID: "42",
	}

	trip, err := svc.CreateTrip(ctx, fare)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("Created trip: %v\n", trip)

	// To prevent the application from exiting immediately, so the tilt file doesnt restart the container again and again
	for {
		time.Sleep(2 * time.Second)
	}
}
