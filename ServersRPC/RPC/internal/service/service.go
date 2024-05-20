package service

import (
	"context"
	"github.com/ekomobile/dadata/v2"
	"github.com/ekomobile/dadata/v2/api/model"
	"github.com/ekomobile/dadata/v2/client"

	"log"
)

type SearchRequest struct {
	Query string `json:"query"`
}

type GeocodeRequest struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type GeoService interface {
	SearchService(request *SearchRequest, reply []*model.Address) error
	GeocodeAddressService(cords string, reply *string) error
}

type Geocoder struct{}

func (g *Geocoder) SearchService(request SearchRequest, reply *[]*model.Address) error {
	log.Println(reply)
	cleanApi := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    "e6b91900da8a4f3c5138bc921a882ee75d42922a",
		SecretKeyValue: "943062a0ae098458484fa91f7947fd31c3f549df"}))
	addresses, err := cleanApi.Address(context.Background(), request.Query)
	if err != nil {
		return err
	}

	*reply = addresses
	return nil
}

func (g *Geocoder) GeocodeAddressService(request GeocodeRequest, reply *[]*model.Address) error {
	cleanApi := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    "e6b91900da8a4f3c5138bc921a882ee75d42922a",
		SecretKeyValue: "943062a0ae098458484fa91f7947fd31c3f549df"}))
	addresses, err := cleanApi.Address(context.Background(), request.Lat, request.Lon)
	if err != nil {
		return err
	}
	*reply = addresses
	return nil
}
