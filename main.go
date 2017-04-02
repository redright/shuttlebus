package main

import (
	"fmt"

	"github.com/redright/shuttlebus/db"
	"github.com/redright/shuttlebus/domain"
)

func main() {
	fmt.Println("Runnig..")
	run := true
	for run {
		cmd := 0
		fmt.Print("Command: ")
		fmt.Scanf("%d\n", &cmd)
		switch cmd {
		case 1:
			var rows = db.Query("select * from passenger")
			fmt.Println(rows)
			for rows.Next() {
				var p = new(domain.Passenger)
				rows.Scan(&p.ID, &p.Name)
				fmt.Println(p)
			}
		case 2:
			var pName string
			fmt.Print("New Passenger Name: ")
			fmt.Scanf("%s\n", &pName)
			r, e := db.Execute("insert into passenger (name) VALUES (?)", pName)
			if e != nil {
				fmt.Println(e)
			}
			fmt.Println(r)
		}

	}

}
