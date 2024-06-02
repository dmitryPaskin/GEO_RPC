package Test

import (
	"GeoServiseAppDate/internal/models"
	"GeoServiseAppDate/internal/service"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestService_Address(t *testing.T) {
	mockResponse := []*models.AddressSearch{
		{Source: "мск сухонска 11/-89",
			Result:               "г Москва, ул Сухонская, д 11, кв 89",
			PostalCode:           "127642",
			Country:              "Россия",
			CountryISOCode:       "RU",
			FederalDistrict:      "Центральный",
			RegionFiasID:         "0c5b2444-70a0-4932-980c-b4dc0d3f02b5",
			RegionKladrID:        "7700000000000",
			RegionISOCode:        "RU-MOW",
			RegionWithType:       "г Москва",
			RegionType:           "г",
			RegionTypeFull:       "город",
			Region:               "Москва",
			CityArea:             "Северо-восточный",
			CityDistrictWithType: "р-н Северное Медведково",
			CityDistrictType:     "р-н",
			CityDistrictTypeFull: "район",
			CityDistrict:         "Северное Медведково",
			SteadFiasID:          "95dbf7fb-0dd4-4a04-8100-4f6c847564b5",
			SteadKladrID:         "77000000000283600",
			StreetWithType:       "ул Сухонская",
			StreetType:           "ул",
			SteadTypeFull:        "улица",
			Street:               "Сухонская",
			HouseFiasID:          "5ee84ac0-eb9a-4b42-b814-2f5f7c27c255",
			HouseKladrID:         "7700000000028360004",
			HouseCadnum:          "77:02:0004008:1017",
			HouseType:            "д",
			HouseTypeFull:        "дом",
			House:                "11",
			FlatFiasID:           "f26b876b-6857-4951-b060-ec6559f04a9a",
			FlatCadnum:           "77:02:0004008:4143",
			FlatType:             "кв",
			FlatTypeFull:         "квартира",
			Flat:                 "89",
			FlatArea:             "34.6",
			SquareMeterPrice:     "244503",
			FlatPrice:            "8459804",
			FiasID:               "f26b876b-6857-4951-b060-ec6559f04a9a",
			FiasCode:             "77000000000000028360004",
			FiasLevel:            "9",
			FiasActualityState:   "0",
			KladrID:              "7700000000028360004",
			CapitalMarker:        "0",
			Okato:                "45280583000",
			Oktmo:                "45362000",
			TaxOffice:            "7715",
			TaxOfficeLegal:       "7715",
			Timezone:             "UTC+3",
			GeoLat:               "55.8782557",
			GeoLon:               "37.65372",
			BeltwayHit:           "IN_MKAD",
			QcGeo:                0,
			QcComplete:           0,
			QcHouse:              2,
			Qc:                   0,
		},
	}

	responseBody, _ := json.Marshal(mockResponse)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, service.UrlAddress, r.URL.String())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	}))
	defer ts.Close()

	client := &http.Client{Timeout: 10 * time.Second}
	service := service.NewService(client)

	request := models.SearchRequest{Query: "мск сухонска 11/-89"}
	result, err := service.Address(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockResponse, result)
}

func TestService_Geocode(t *testing.T) {
	mockResponse := &models.AddressGeo{ /* заполняем данными */ }
	responseBody, _ := json.Marshal(mockResponse)

	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "POST", r.Method)
		assert.Equal(t, service.UrlGeocode, r.URL.String())
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(responseBody)
	}))
	defer ts.Close()

	client := &http.Client{Timeout: 10 * time.Second}
	service := service.NewService(client)

	request := models.GeocodeRequest{ /* заполняем данными */ }
	result, err := service.Geocode(request)

	assert.NoError(t, err)
	assert.NotNil(t, result)
	assert.Equal(t, mockResponse, result)
}
