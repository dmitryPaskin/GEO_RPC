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
		ApiKeyValue:    "9a84b6e525fb548e7170b77175e9e15af84a30ac",
		SecretKeyValue: "6ecfe8510311d14daf5de31de9a5af4ceeb5b0d5"}))
	addresses, err := cleanApi.Address(context.Background(), request.Query)
	if err != nil {
		return err
	}

	*reply = addresses
	return nil
}

func (g *Geocoder) GeocodeAddressService(request GeocodeRequest, reply *[]*model.Address) error {
	cleanApi := dadata.NewCleanApi(client.WithCredentialProvider(&client.Credentials{
		ApiKeyValue:    "9a84b6e525fb548e7170b77175e9e15af84a30ac",
		SecretKeyValue: "6ecfe8510311d14daf5de31de9a5af4ceeb5b0d5"}))
	addresses, err := cleanApi.Address(context.Background(), request.Lat, request.Lon)
	if err != nil {
		return err
	}
	*reply = addresses
	return nil
}
