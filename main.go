package main

import (
	"fmt"

	"github.com/redright/shuttlebus/passenger"
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
			r := *passenger.GetPassegers()
			for i := 0; i < len(r); i++ {
				fmt.Println(r[i])
			}
		case 2:
			var newPassenger passenger.Passenger
			fmt.Println("New Passenger")
			fmt.Print(" Name: ")
			fmt.Scanf("%s\n", &newPassenger.Name)
			fmt.Print(" Last Name: ")
			fmt.Scanf("%s\n", &newPassenger.LastName)
			fmt.Print(" Phone Number: ")
			fmt.Scanf("%s\n", &newPassenger.PhoneNumber)
			passenger.Create(&newPassenger)
			fmt.Println(newPassenger)
		}
	}

}
