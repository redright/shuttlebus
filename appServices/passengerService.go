package appServices

import (
	"github.com/redright/shuttlebus/db"
	"github.com/redright/shuttlebus/domain"
)

type PassengerService struct {
	repo db.PassengerRepo
}

func (t *PassengerService) Init() {
	t.repo = db.PassengerRepo{}
}

func (t *PassengerService) GetPassengers() *[]domain.Passenger {
	return t.repo.GetPassegers()
}
func (t *PassengerService) Create(p *domain.Passenger) *domain.Passenger {
	t.repo.Create(p)
	return p
}
