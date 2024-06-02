package Test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"GeoServiseAppDate/internal/controller/responder"
	handler "GeoServiseAppDate/internal/controller/searchGEOHandlers"
	"GeoServiseAppDate/internal/models"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockService struct {
	mock.Mock
}

func (m *MockService) Address(request models.SearchRequest) ([]*models.AddressSearch, error) {
	args := m.Called(request)
	return args.Get(0).([]*models.AddressSearch), args.Error(1)
}

func (m *MockService) Geocode(request models.GeocodeRequest) (*models.AddressGeo, error) {
	args := m.Called(request)
	return args.Get(0).(*models.AddressGeo), args.Error(1)
}

type MockResponder struct {
	mock.Mock
}

func (m *MockResponder) ErrorBedRequest(w http.ResponseWriter, err error) {
	m.Called(w, err)
	w.WriteHeader(http.StatusBadRequest)
}

func (m *MockResponder) ErrorInternal(w http.ResponseWriter, err error) {
	m.Called(w, err)
	w.WriteHeader(http.StatusInternalServerError)
}

func (m *MockResponder) OutputJSON(w http.ResponseWriter, response responder.Response) {
	m.Called(w, response)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func TestSearchAddressHandler(t *testing.T) {
	mockService := new(MockService)
	mockResponder := new(MockResponder)
	handler := handler.New(nil, nil)

	searchRequest := models.SearchRequest{Query: "some address"}
	expectedResponse := []*models.AddressSearch{
		{ /* заполняем данными */ },
	}
	mockService.On("Address", searchRequest).Return(expectedResponse, nil)
	mockResponder.On("OutputJSON", mock.Anything, responder.Response{
		Success: true,
		Message: "address get",
		Data:    expectedResponse,
	}).Return()

	reqBody, _ := json.Marshal(searchRequest)
	req, err := http.NewRequest("POST", "/address/search", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.GeocodeHandler)
	handlerFunc.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
	mockResponder.AssertExpectations(t)
}

func TestGeocodeHandler(t *testing.T) {
	mockService := new(MockService)
	mockResponder := new(MockResponder)
	handler := handler.New(nil, nil)

	geocodeRequest := models.GeocodeRequest{ /* заполняем данными */ }
	expectedResponse := &models.AddressGeo{ /* заполняем данными */ }
	mockService.On("Geocode", geocodeRequest).Return(expectedResponse, nil)
	mockResponder.On("OutputJSON", mock.Anything, responder.Response{
		Success: true,
		Message: "address get",
		Data:    expectedResponse,
	}).Return()

	reqBody, _ := json.Marshal(geocodeRequest)
	req, err := http.NewRequest("POST", "/address/geocode", bytes.NewBuffer(reqBody))
	assert.NoError(t, err)

	rr := httptest.NewRecorder()
	handlerFunc := http.HandlerFunc(handler.GeocodeHandler)
	handlerFunc.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	mockService.AssertExpectations(t)
	mockResponder.AssertExpectations(t)
}
