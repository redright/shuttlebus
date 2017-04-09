package db

import (
	"github.com/redright/shuttlebus/domain"
)

type UserRepo struct {
}

func (r *UserRepo) CreateUser(user *domain.User) {
	user.ID = GenerateID()
	Execute("INSERT INTO user (ID,Name,LastName,Email) VALUES (?,?,?,?)", user.ID, user.Name, user.LastName, user.Email)
}

func (r *UserRepo) DeleteUser(id string) {
	deleteRow("user", id)
}

func (r *UserRepo) CreateUserPhone(phone *domain.UserPhone) {
	phone.ID = GenerateID()
	Execute("INSERT INTO userphone (ID,UserID,AreaCode,Number) VALUES (?,?,?,?)",
		phone.ID, phone.UserID, phone.AreaCode, phone.Number)
}

func (r *UserRepo) DeleteUserPhone(id string) {
	deleteRow("userphone", id)
}
