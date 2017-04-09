package domain

type Client struct {
	ID   string
	Name string
}

type ClientPassenger struct {
	User
	ID       string
	UserID   string
	ClientID string
}
