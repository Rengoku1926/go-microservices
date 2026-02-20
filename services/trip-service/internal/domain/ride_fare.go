package domain


import (
	pb "ride-sharing/shared/proto/trip"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RideFareModel struct {
	ID primitive.ObjectID
	UserId string
	PackageSlug string
	TotalPriceInCents float64
}


func (r *RideFareModel) ToProto() *pb.RideFare {
	return &pb.RideFare{
		Id:                r.ID.Hex(),
		UserID:            r.UserId,
		PackageSlug:       r.PackageSlug,
		TotalPriceInCents: r.TotalPriceInCents,
	}
}

func ToRideFaresProto(fares []*RideFareModel) []*pb.RideFare {
	var protoFares []*pb.RideFare
	for _, f := range fares {
		protoFares = append(protoFares, f.ToProto())
	}
	return protoFares
}