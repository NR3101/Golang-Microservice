package types

import (
	pb "ride-sharing/shared/proto/trip"
)

type OSRMApiResponse struct {
	Routes []struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Geometry struct {
			Coordinates [][]float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"routes"`
}

func (r *OSRMApiResponse) ToProto() *pb.Route {
	if len(r.Routes) == 0 {
		return &pb.Route{}
	}

	route := r.Routes[0]
	geometry := route.Geometry.Coordinates
	coordinates := make([]*pb.Coordinate, len(geometry))
	for i, coord := range geometry {
		coordinates[i] = &pb.Coordinate{
			Latitude:  coord[1],
			Longitude: coord[0],
		}
	}

	grpcRoute := &pb.Route{
		Geometry: []*pb.Geometry{
			{
				Coordinates: coordinates,
			},
		},
		Distance: route.Distance,
		Duration: route.Duration,
	}

	return grpcRoute
}
