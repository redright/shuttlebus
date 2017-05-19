package appServices

import (
	"time"

	"github.com/redright/shuttlebus/db"
	"github.com/redright/shuttlebus/domain"
)

type PassengerService struct {
	BaseService
}

func (s *PassengerService) ShareLocationToShuttle(shuttleID string, location domain.Point) {
	//TODO: check shuttleid
	shuttleRepo := db.ShuttleRepo{}
	shuttleRepo.AddPassengerLocation(domain.PassengerLocation{
		ShuttleID:   shuttleID,
		PassengerID: s.Context.PassengerID,
		Location:    location,
		SharingTime: time.Now(),
	})
}

func (s *PassengerService) GetShuttles() []domain.Shuttle {
	shuttleRepo := db.ShuttleRepo{}
	return shuttleRepo.GetShuttles(s.Context.PassengerID)
}
