package db

import "github.com/redright/shuttlebus/domain"

type PassengerRepo struct {
}

func (r *PassengerRepo) Create(p *domain.Passenger) {
	if p.ID > 0 {
		panic("PassengerAlreadyCreated")
	}

	res := Execute("insert into passenger (Name,LastName,PhoneNumber) VALUES(?,?,?)", &p.Name, &p.LastName, &p.PhoneNumber)
	p.ID, _ = res.LastInsertId()
}

func (r *PassengerRepo) GetPassenger(id int64) *domain.Passenger {
	return nil
	// rows := Query("select ID,Name,LastName,PhoneNumber where ID=?", &id)
	// result := make([]Passenger, 0)
	// for rows.Next() {
	// 	var p = &Passenger
	// 	append(result, p)
	// }
	// return &result
}

func (r *PassengerRepo) GetPassegers() *[]domain.Passenger {
	rows := Query("select ID,Name,LastName,PhoneNumber from passenger")
	var result []domain.Passenger
	for rows.Next() {
		var p domain.Passenger
		rows.Scan(&p.ID, &p.Name, &p.LastName, &p.PhoneNumber)
		result = append(result, p)
	}
	return &result
}
