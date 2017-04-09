package domain

type User struct {
	ID           string
	Name         string
	LastName     string
	Email        string
	PhoneNumbers []UserPhone
}

type UserPhone struct {
	ID       string
	UserID   string
	AreaCode string
	Number   string
}
