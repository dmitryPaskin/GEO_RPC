package service

import (
	proto "GeoServiseAppDate/ServersRPC/gRPC/pkg/GEOgRPC/api/GeoService"
	"context"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/client"
	"github.com/mitchellh/mapstructure"
)

type Geocoder struct {
	proto.UnimplementedGeoServiceServer
}

func (g *Geocoder) SearchService(ctx context.Context, in *proto.SearchRequest) (*proto.AddressSearch, error) {
	var (
		result proto.AddressSearch
		err    error
	)
	cleanApi := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    "e6b91900da8a4f3c5138bc921a882ee75d42922a",
		SecretKeyValue: "943062a0ae098458484fa91f7947fd31c3f549df"}))

	addresses, err := cleanApi.Address(context.Background(), in.Query)
	if err != nil {
		return nil, err
	}

	if err = mapstructure.Decode(addresses, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

func (g *Geocoder) GeocodeAddressService(ctx context.Context, in *proto.GeocodeRequest) (*proto.AddressGeo, error) {
	var (
		result proto.AddressGeo
		err    error
	)

	cleanApi := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    "e6b91900da8a4f3c5138bc921a882ee75d42922a",
		SecretKeyValue: "943062a0ae098458484fa91f7947fd31c3f549df"}))
	addresses, err := cleanApi.Address(context.Background(), in.Lat, in.Lon)
	if err != nil {
		return nil, err
	}

	if err = mapstructure.Decode(addresses, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
