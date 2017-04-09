package main

import (
	"github.com/redright/shuttlebus/APIHost"
)

func main() {
	APIHost.Setup()
	// run := true
	// var repo = new(db.PassengerRepo)
	// for run {
	// 	cmd := 0
	// 	fmt.Print("Command: ")
	// 	fmt.Scanf("%d\n", &cmd)
	// 	switch cmd {
	// 	case 1:
	// 		r := *repo.GetPassegers()
	// 		for i := 0; i < len(r); i++ {
	// 			fmt.Println(r[i])
	// 		}
	// 	case 2:
	// 		var newPassenger domain.Passenger
	// 		fmt.Println("New Passenger")
	// 		fmt.Print(" Name: ")
	// 		fmt.Scanf("%s\n", &newPassenger.Name)
	// 		fmt.Print(" Last Name: ")
	// 		fmt.Scanf("%s\n", &newPassenger.LastName)
	// 		fmt.Print(" Phone Number: ")
	// 		fmt.Scanf("%s\n", &newPassenger.PhoneNumber)
	// 		repo.Create(&newPassenger)
	// 		fmt.Println(newPassenger)
	// 	}
	// }

}
