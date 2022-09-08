package tripservice

import (
	"context"
	trippb "coolcar/gen/go/trip"
)

// type TripServiceServer interface {
// 	GetTrip(context.Context, *GetTripRequest) (*GetTripResponse, error)
// 	mustEmbedUnimplementedTripServiceServer()
// }

// Service
type Service struct {
	// 内嵌一个内置的UnimplementedTripServiceServer来实现TripServiceServer接口
	trippb.UnimplementedTripServiceServer
}

func (*Service) GetTrip(c context.Context, req *trippb.GetTripRequest) (*trippb.GetTripResponse, error) {
	return &trippb.GetTripResponse{
		Id: req.Id,
		Trip: &trippb.Trip{
			Start:       "abc",
			End:         "cda",
			DurationSec: 1000,
			FeeCent:     1000,
			StartPos: &trippb.Location{
				Latitude:  30,
				Longitude: 120,
			},
			EndPos: &trippb.Location{
				Latitude:  35,
				Longitude: 115,
			},
			PathLocations: []*trippb.Location{
				{
					Latitude:  110,
					Longitude: 110,
				},
				{
					Latitude:  100,
					Longitude: 100,
				},
			},
			Status: trippb.TripStatus_FINISHED,
		},
	}, nil
}
