package db

import (
	"github.com/redright/shuttlebus/domain"
)

type ClientRepo struct {
}

func (r *ClientRepo) CreateClient(client *domain.Client) {
	client.ID = GenerateID()
	Execute("INSERT INTO client (ID,Name) VALUES (?,?)", client.ID, client.Name)
}

func (r *ClientRepo) DeleteClient(id string) {
	deleteRow("client", id)
}

func (r *ClientRepo) CreatePassenger(passenger *domain.ClientPassenger) {
	passenger.ID = GenerateID()
	if passenger.UserID == "" {
		var userRepo UserRepo
		userRepo.CreateUser(&passenger.User)
	}
	Execute("INSERT INTO clientPassenger (ID,UserID,ClientID) VALUES(?,?,?)", passenger.ID, passenger.User.ID, passenger.ClientID)
}
