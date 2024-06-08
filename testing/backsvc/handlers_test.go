package service

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

var (
	formatter = render.New(render.Options{
		IndentJSON: true,
	})
)

// Dummy fulfillmentStatus struct for testing
type fulfillmentStatus struct {
	QuantityInStock int    `json:"quantity_in_stock"`
	ShipsWithin     int    `json:"ships_within"`
	SKU             string `json:"sku"`
}

// NewServer sets up the test server
func NewServer() http.Handler {
	router := mux.NewRouter()
	router.HandleFunc("/skus/{sku}", getFulfillmentStatus).Methods("GET")
	n := negroni.Classic()
	n.UseHandler(router)
	return n
}

// getFulfillmentStatus is the handler function for the test
func getFulfillmentStatus(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sku := params["sku"]
	// Mock response
	if sku == "THINGAMAJIG12" {
		response := fulfillmentStatus{
			QuantityInStock: 100,
			ShipsWithin:     14,
			SKU:             sku,
		}
		formatter.JSON(w, http.StatusOK, response)
	} else {
		http.NotFound(w, r)
	}
}

func TestGetFullfillmentStatusReturns200ForExistingSKU(t *testing.T) {
	var (
		request  *http.Request
		recorder *httptest.ResponseRecorder
	)

	server := NewServer()
	targetSKU := "THINGAMAJIG12"

	recorder = httptest.NewRecorder()
	request, _ = http.NewRequest("GET", "/skus/"+targetSKU, nil)
	server.ServeHTTP(recorder, request)

	var detail fulfillmentStatus

	if recorder.Code != http.StatusOK {
		t.Errorf("Success expected: %d", recorder.Code)
	}
	payload, err := io.ReadAll(recorder.Body)
	if err != nil {
		t.Error(err)
	}
	err = json.Unmarshal(payload, &detail)
	if err != nil {
		t.Errorf("Error unmarshalling JSON: %s", err)
	}
	if detail.QuantityInStock != 100 {
		t.Errorf("Expected 100, got %d", detail.QuantityInStock)
	}
	if detail.ShipsWithin != 14 {
		t.Errorf("Expected 14, got %d", detail.ShipsWithin)
	}
	if detail.SKU != "THINGAMAJIG12" {
		t.Errorf("Expected THINGAMAJIG12, got %s", detail.SKU)
	}
}
