package domain

import (
	pb "ride-sharing/shared/proto/trip"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID                primitive.ObjectID
	UserID            string
	PackageSlug       string // e.g. "van", "sedan", etc.
	TotalPriceInCents float64
	ExpiresAt         time.Time
}

func (r *RideFareModel) ToProto() *pb.RideFare {
	return &pb.RideFare{
		Id:                r.ID.Hex(),
		UserID:            r.UserID,
		PackageSlug:       r.PackageSlug,
		TotalPriceInCents: r.TotalPriceInCents,
	}
}

func ToRideFaresProto(fares []*RideFareModel) []*pb.RideFare {
	protoFares := make([]*pb.RideFare, len(fares))
	for i, fare := range fares {
		protoFares[i] = fare.ToProto()
	}
	return protoFares
}
