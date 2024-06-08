package service

import (
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"net/http"
)

func getFullfillmentStatusHandler(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		sku := vars["sku"]
		formatter.JSON(w, http.StatusOK, fulfillmentStatus{
			SKU:             sku,
			ShipsWithin:     14,
			QuantityInStock: 100,
		})
	}
}

func rootHandel(formatter *render.Render) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		formatter.Text(w, http.StatusOK, "Fulfillment service, see (url) for API")
	}
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/skus/{sku}",
		getFullfillmentStatusHandler(formatter)).Methods("GET")
}
