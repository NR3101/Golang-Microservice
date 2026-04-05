package grpc

import (
	"context"
	"ride-sharing/services/trip-service/internal/domain"
	pb "ride-sharing/shared/proto/trip"
	"ride-sharing/shared/types"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type gRPCHandler struct {
	pb.UnimplementedTripServiceServer
	service domain.TripService
}

func NewGRPCHandler(server *grpc.Server, service domain.TripService) *gRPCHandler {
	handler := &gRPCHandler{service: service}
	pb.RegisterTripServiceServer(server, handler)
	return handler
}

func (h *gRPCHandler) PreviewTrip(ctx context.Context, req *pb.PreviewTripRequest) (*pb.PreviewTripResponse, error) {
	pickUp := req.GetStartLocation()
	destination := req.GetEndLocation()

	pickupCoords := &types.Coordinate{
		Latitude:  pickUp.GetLatitude(),
		Longitude: pickUp.GetLongitude(),
	}
	destinationCoords := &types.Coordinate{
		Latitude:  destination.GetLatitude(),
		Longitude: destination.GetLongitude(),
	}

	userID := req.GetUserID()

	t, err := h.service.GetRoute(ctx, pickupCoords, destinationCoords)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to get route: %v", err)
	}

	// Estimate fares based on the route
	estimatedFares := h.service.EstimatePackagesPriceWithRoute(t)

	// Save the estimated fares to the database
	fares, err := h.service.GenerateTripFares(ctx, estimatedFares, userID)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to save ride fares: %v", err)
	}

	return &pb.PreviewTripResponse{
		Route:     t.ToProto(),
		RideFares: domain.ToRideFaresProto(fares),
	}, nil
}

func (h *gRPCHandler) CreateTrip(ctx context.Context, req *pb.CreateTripRequest) (*pb.CreateTripResponse, error) {
	return nil, status.Error(codes.Unimplemented, "method CreateTrip not implemented")
}
