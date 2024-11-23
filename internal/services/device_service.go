package services

import (
	"errors"
	"github.com/ShaharGutmanCoding/SmartHome/internal/models"
)

var devices = []models.Device{
	{ID: 1, Name: "Living Room Light", Type: "Light", Status: "Off"},
	{ID: 2, Name: "Bedroom Fan", Type: "Fan", Status: "On"},
}

// GetAllDevices retrieves all devices
func GetAllDevices() []models.Device {
	return devices
}

// GetDeviceByID retrieves a device by its ID
func GetDeviceByID(id int) (models.Device, error) {
	for _, device := range devices {
		if device.ID == id {
			return device, nil
		}
	}
	return models.Device{}, errors.New("device not found")
}

// CreateDevice creates a new device
func CreateDevice(newDevice models.Device) models.Device {
	newDevice.ID = len(devices) + 1
	devices = append(devices, newDevice)
	return newDevice
}

// UpdateDevice updates an existing device
func UpdateDevice(id int, updatedDevice models.Device) (models.Device, error) {
	for i, device := range devices {
		if device.ID == id {
			updatedDevice.ID = id
			devices[i] = updatedDevice
			return updatedDevice, nil
		}
	}
	return models.Device{}, errors.New("device not found")
}

// DeleteDevice deletes a device by ID
func DeleteDevice(id int) error {
	for i, device := range devices {
		if device.ID == id {
			devices = append(devices[:i], devices[i+1:]...)
			return nil
		}
	}
	return errors.New("device not found")
}
