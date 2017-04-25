package domain

import (
	"time"
)

type Shuttle struct {
	ID               string
	CompanyID        string
	ClientID         string
	CompanyVehicleID string
	CompanyDriverID  string

	// Company    Company
	// Client     Client
	// Vehicle    Vehicle
	StartPoint Point
	EndPoint   Point
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
