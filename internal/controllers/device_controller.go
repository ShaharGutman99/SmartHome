package controllers

import (
	"encoding/json"
	"net/http"
	"github.com/ShaharGutmanCoding/SmartHome/internal/models"
	"github.com/ShaharGutmanCoding/SmartHome/internal/services"
	"strconv"

	"github.com/gorilla/mux"
)

// GetAllDevices handles GET /devices
func GetAllDevices(w http.ResponseWriter, r *http.Request) {
	devices := services.GetAllDevices()
	json.NewEncoder(w).Encode(devices)
}

// GetDeviceByID handles GET /devices/{id}
func GetDeviceByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	device, err := services.GetDeviceByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(device)
}

// CreateDevice handles POST /devices
func CreateDevice(w http.ResponseWriter, r *http.Request) {
	var newDevice models.Device
	if err := json.NewDecoder(r.Body).Decode(&newDevice); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	device := services.CreateDevice(newDevice)
	json.NewEncoder(w).Encode(device)
}

// UpdateDevice handles PUT /devices/{id}
func UpdateDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedDevice models.Device
	if err := json.NewDecoder(r.Body).Decode(&updatedDevice); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	device, err := services.UpdateDevice(id, updatedDevice)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(device)
}

// DeleteDevice handles DELETE /devices/{id}
func DeleteDevice(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := services.DeleteDevice(id); err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
