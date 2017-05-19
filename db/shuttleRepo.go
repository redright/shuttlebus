package db

import (
	"github.com/redright/shuttlebus/domain"
)

type ShuttleRepo struct {
}

func (r *ShuttleRepo) CreateShuttle(shuttle *domain.Shuttle) {
	shuttle.ID = GenerateID()
	Execute("INSERT INTO Shuttle (ID,CompanyID,ClientID,CompanyVehicleID,CompanyDriverID,StartLat,StartLng,EndLat,EndLng VALUES (?,?,?,?,?,?,?,?,?)",
		shuttle.ID, shuttle.CompanyID, shuttle.ClientID, shuttle.CompanyVehicleID, shuttle.CompanyDriverID, shuttle.From.Latitude, shuttle.From.Longitute, shuttle.To.Latitude, shuttle.To.Longitute)
}

func (r ShuttleRepo) AddPassenger(shuttle *domain.Shuttle, passenger *domain.ClientPassenger) {
	Execute("INSERT INTO ShuttlePassenger (PassengerID,ShuttleID) VALUES (?,?)", shuttle.ID, passenger.ID)
}

func (r ShuttleRepo) GetShuttles(passengerID string) []domain.Shuttle {
	rows := Query(`select s.ID,s.CompanyID,s.CompanyDriverID,s.CompanyVehicleID,s.FromLat,s.FromLng,s.ToLat,s.ToLng
,s.Name 
,s.Label
,v.PlateNo
from shuttle s
inner join shuttlepassenger sp
on sp.ShuttleID = s.ID
left join companyvehicle cv
on cv.ID = s.CompanyVehicleID
left join vehicle v
on v.ID = cv.VehicleID	
where sp.PassengerID = ?`, passengerID)
	var result []domain.Shuttle
	for rows.Next() {
		var row domain.Shuttle
		rows.Scan(&row.ID,
			&row.CompanyID,
			&row.CompanyDriverID,
			&row.CompanyVehicleID,
			&row.From.Latitude,
			&row.From.Longitute,
			&row.To.Latitude,
			&row.To.Longitute,
			//&row.ClientID,
			&row.Name,
			&row.Label,
			&row.Vehicle.PlateNo)
		result = append(result, row)
	}
	return result
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
