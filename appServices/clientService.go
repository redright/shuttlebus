package appServices

import (
	"github.com/redright/shuttlebus/common"
	"github.com/redright/shuttlebus/db"
	"github.com/redright/shuttlebus/domain"
)

type ClientService struct {
	repo db.ClientRepo
}

func (s *ClientService) CreateClient(c domain.Client) (string, error) {
	if len(c.Name) == 0 {
		return "", &common.BusinessError{Code: "InvalidClientName"}
	}
	s.repo.CreateClient(&c)
	return "Başarıyla Oluşturuldu", nil
}

func (s *ClientService) AddPassenger(p domain.ClientPassenger) {
	s.repo.CreatePassenger(&p)
}
