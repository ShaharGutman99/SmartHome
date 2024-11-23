package routes

import (
	"github.com/ShaharGutmanCoding/SmartHome/internal/controllers"
	"github.com/gorilla/mux"
)

// SetupRoutes initializes the API routes
func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	// API versioning
	api := router.PathPrefix("/api/v1").Subrouter()

	// Device routes
	api.HandleFunc("/devices", controllers.GetAllDevices).Methods("GET")
	api.HandleFunc("/devices/{id}", controllers.GetDeviceByID).Methods("GET")
	api.HandleFunc("/devices", controllers.CreateDevice).Methods("POST")
	api.HandleFunc("/devices/{id}", controllers.UpdateDevice).Methods("PUT")
	api.HandleFunc("/devices/{id}", controllers.DeleteDevice).Methods("DELETE")

	return router
}
