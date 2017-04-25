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

func (r ShuttleRepo) GetShuttles(passengerID string) []domain.Shuttle {
	return nil
}

func (r ShuttleRepo) AddPassengerLocation(pLocation domain.PassengerLocation) {
	Execute("INSERT INTO ShuttlePassengerLocation (ShuttleID,PassengerID,Lat,Lng,SharingTime) VALUES(?,?,?,?,?)",
		pLocation.ShuttleID,
		pLocation.PassengerID,
		pLocation.Location.Latitude,
		pLocation.Location.Longitute,
		pLocation.SharingTime)
}

func (r ShuttleRepo) AddShuttleLocation(sLocation domain.ShuttleLocation) {
	Execute("INSERT INTO ShuttleLocation (ShuttleID,Lat,Lng,SharingTime) VALUES (?,?,?,?)",
		sLocation.ShuttleID,
		sLocation.Location.Latitude,
		sLocation.Location.Longitute,
		sLocation.SharingTime)
}
