package domain

import (
	"time"
)

type Company struct {
	ID      string
	Name    string
	Clients []Client
}

type CompanyDriver struct {
	Driver
	ID        string
	DriverID  string
	CompanyID string
	HireDate  time.Time
	Status    string
}

type CompanyVehicle struct {
	Vehicle
	ID        string
	CompanyID string
	VehicleID string
}
