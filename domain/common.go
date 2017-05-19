package domain

import (
	"time"
)

type Driver struct {
	User
	ID               string
	UserID           string
	LicenceNumber    string
	LicenceTakenDate time.Time
}

type Vehicle struct {
	ID        string
	CompanyID string
	PlateNo   string `json:"plateNo"`
	Capacity  int
	BrandCode string
	ModelCode string
	ModelYear int
}

type Point struct {
	Latitude  float32 `json:"latitude"`
	Longitute float32 `json:"longitute"`
}
