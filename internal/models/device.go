package models

// Device represents a smart home device
type Device struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"` // e.g., "On", "Off"
}
