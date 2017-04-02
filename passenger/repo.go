package passenger

import "github.com/redright/shuttlebus/db"

func Create(p *Passenger) {
	if p.ID > 0 {
		panic("PassengerAlreadyCreated")
	}

	r := db.Execute("insert into passenger (Name,LastName,PhoneNumber) VALUES(?,?,?)", &p.Name, &p.LastName, &p.PhoneNumber)
	p.ID, _ = r.LastInsertId()
}

func GetPassenger(id int64) *Passenger {
	return nil
	// rows := db.Query("select ID,Name,LastName,PhoneNumber where ID=?", &id)
	// result := make([]Passenger, 0)
	// for rows.Next() {
	// 	var p = &Passenger
	// 	append(result, p)
	// }
	// return &result
}

func GetPassegers() *[]Passenger {
	rows := db.Query("select ID,Name,LastName,PhoneNumber from passenger")
	var result []Passenger
	for rows.Next() {
		var p Passenger
		rows.Scan(&p.ID, &p.Name, &p.LastName, &p.PhoneNumber)
		result = append(result, p)
	}
	return &result
}
