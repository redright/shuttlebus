package domain

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
