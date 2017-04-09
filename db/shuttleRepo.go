package db

import (
	"github.com/redright/shuttlebus/domain"
)

type ShuttleRepo struct {
}

func (r *ShuttleRepo) CreateShuttle(shuttle *domain.Shuttle) {
	shuttle.ID = GenerateID()
	Execute("INSERT INTO Shuttle (ID,CompanyID,ClientID,CompanyVehicleID,CompanyDriverID,StartLat,StartLng,EndLat,EndLng VALUES (?,?,?,?,?,?,?,?,?)",
		shuttle.ID, shuttle.CompanyID, shuttle.ClientID, shuttle.CompanyVehicleID, shuttle.CompanyDriverID, shuttle.StartPoint.Latitude, shuttle.StartPoint.Longitute, shuttle.EndPoint.Latitude, shuttle.EndPoint.Longitute)
}

func (r ShuttleRepo) AddPassenger(shuttle *domain.Shuttle, passenger *domain.ClientPassenger) {
	Execute("INSERT INTO ShuttlePassenger (PassengerID,ShuttleID) VALUES (?,?)", shuttle.ID, passenger.ID)
}
