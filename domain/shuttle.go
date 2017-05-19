package domain

import (
	"time"
)

type Shuttle struct {
	ID               string `json:"id"`
	CompanyID        string `json:"companyID"`
	ClientID         string `json:"clientID"`
	CompanyVehicleID string `json:"companyVehicleID"`
	CompanyDriverID  string `json:"companyDriverID"`

	// Company    Company
	// Client     Client
	Vehicle Vehicle `json:"vehicle"`
	From    Point   `json:"from"`
	To      Point   `json:"to"`
	Name    string  `json:"name"`
	Label   string  `json:"label"`
	// Passengers []ClientPassenger
}

type PassengerLocation struct {
	ShuttleID   string
	PassengerID string
	Location    Point
	SharingTime time.Time
}
type ShuttleLocation struct {
	ShuttleID   string
	Location    Point
	SharingTime time.Time
}
