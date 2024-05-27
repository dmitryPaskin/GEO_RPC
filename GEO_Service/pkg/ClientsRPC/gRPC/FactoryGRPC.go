package gRPC

import (
	"GeoServiseAppDate/internal/models"
	pb "GeoServiseAppDate/pkg/ClientsRPC/gRPC/pkg/GEOgRPC/proto"
	"context"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc"
)

type GRPC struct {
}

func (g *GRPC) Address(request models.SearchRequest) ([]*models.AddressSearch, error) {
	var result []*models.AddressSearch
	conn, err := grpc.Dial("grpc:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewGeoServiceClient(conn)
	req := &pb.SearchRequest{Query: request.Query}
	res, err := client.SearchService(context.Background(), req)
	if err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(res, &result); err != nil {
		return nil, err
	}
	return result, nil
}

func (g *GRPC) Geocode(request models.GeocodeRequest) (*models.AddressGeo, error) {
	var result *models.AddressGeo
	conn, err := grpc.Dial("grpc:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewGeoServiceClient(conn)
	req := &pb.GeocodeRequest{Lon: request.Lon, Lat: request.Lat}
	res, err := client.GeocodeAddressService(context.Background(), req)
	if err != nil {
		return nil, err
	}

	if err := mapstructure.Decode(res, &result); err != nil {
		return nil, err
	}
	return result, nil
}
