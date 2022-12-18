package main

import (
	"Assignment/Users"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

outer:
	for {
		fmt.Println("====================")
		fmt.Println("Ride Sharing\n",
			"1.Access Passenger Menu\n",
			"2.Access Driver Menu\n",
			"9.Quit\n")
		fmt.Println("====================")

		fmt.Print("Enter an option: ")
		var choice string
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(input)

		switch choice {
		case "1":
			Users.PassengerMenu()
		case "2":
			Users.DriverMenu()
		case "9":
			fmt.Println("You are now exiting the app")
			break outer
		}
	}
}
