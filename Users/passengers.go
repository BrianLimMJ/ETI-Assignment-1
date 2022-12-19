package Users

import (
	Trips "Assignment/Trip"
	"bufio"
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

type Passengers struct {
	passengerID int
	firstName   string
	lastName    string
	mobileNo    int
	emailAdd    string
}

var loggedInPassenger Passengers

func PassengerMenu() {
menu:
	for {
		fmt.Println("====================")
		fmt.Println("Passenger Menu\n",
			"1.Create Passenger Account\n",
			"2.Login into Passenger Account\n",
			"9.Quit\n")
		fmt.Println("====================")

		fmt.Print("Enter an option: ")
		var choice string
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(input)

		switch choice {
		case "1":
			CreatePassengerAccount()
		case "2":
			LogInAsPassenger()
		case "3":
			//login()
		case "9":
			break menu
		}
	}
}

// Creating new Passenger account
func CreatePassengerAccount() {
	var newPassenger Passengers

	fmt.Print("Enter First Name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newPassenger.firstName = strings.TrimSpace(input)

	fmt.Print("Enter Last Name: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newPassenger.lastName = strings.TrimSpace(input)

	fmt.Print("Enter Mobile Number: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	x := strings.TrimSpace(input)
	newPassenger.mobileNo, _ = strconv.Atoi(x)

	fmt.Print("Enter Email Address: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newPassenger.emailAdd = strings.TrimSpace(input)

	client := &http.Client{}
	url := "https://localhost:5000/Passenger"
	postBody, _ := json.Marshal(newPassenger)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest("POST", url, resBody); err == nil {
		if res, err2 := client.Do(req); err2 == nil {
			if res.StatusCode == 202 {
				fmt.Println("Passenger account has been successfully created")
			} else if res.StatusCode == 400 {
				fmt.Println("Bad Request")
			}
		}
	}
	// //Calling of database
	// db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

	// // Error handling
	// if err != nil {
	// 	fmt.Println("Error in connecting to database")
	// 	panic(err.Error())
	// }
	// defer db.Close()

	// //Inserting values into database
	// _, err = db.Exec("insert into passengers (firstName, lastName, mobileNo, emailAdd) values(?, ?, ?, ?)",
	// 	newPassenger.firstName, newPassenger.lastName, newPassenger.mobileNo, newPassenger.emailAdd)
	// if err != nil {
	// 	fmt.Println("Error with sending data to database")
	// 	panic(err.Error())

	// } else {
	// 	// To notify of successful account creation
	// 	fmt.Println("====================")
	// 	fmt.Println("Passenger account has been successfully created")
	// }
}

// Logging in as Passenger
func LogInAsPassenger() {
	fmt.Print("Enter your mobile number: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	PassengerMobileNo := strings.TrimSpace(input)

	//Calling of database
	db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

	// Error handling
	if err != nil {
		fmt.Println("Error in connecting to database")
		panic(err.Error())
	}
	defer db.Close()

	//Checking for value in database
	result, err := db.Query("select * from passengers where mobileNo = ?", PassengerMobileNo)
	if err != nil {
		fmt.Println("Error with getting data from database")
		panic(err.Error())

	} else {
		for result.Next() {
			var PassengerInfo Passengers
			err = result.Scan(&PassengerInfo.passengerID, &PassengerInfo.firstName, &PassengerInfo.lastName, &PassengerInfo.mobileNo, &PassengerInfo.emailAdd)
			x := strconv.Itoa(PassengerInfo.mobileNo)
			if err != nil {
				fmt.Println("Unsuccessful Login")
				panic(err.Error())

			} else if x == PassengerMobileNo {
				loggedInPassenger = PassengerInfo
				// To notify of successful login
				fmt.Println("====================")
				fmt.Println("Successfully logged into " + PassengerInfo.firstName)
				// Logged in as Passenger section
			loggedIn:
				for {
					fmt.Print("\n")
					fmt.Println("====================")
					fmt.Println("You are now logged in as a Passenger\n",
						"1.Update Passenger Account\n",
						"2.Request Trip\n",
						"3.Retrieve Trips\n",
						"9.Quit\n")
					fmt.Println("====================")

					fmt.Print("Enter an option: ")
					var choice string
					reader := bufio.NewReader(os.Stdin)
					input, _ := reader.ReadString('\n')
					choice = strings.TrimSpace(input)

					switch choice {
					case "1":
						UpdatePassengerAccount()
					case "2":
						Trips.RequestTrip(loggedInPassenger.passengerID)
					case "3":
						Trips.RetriveTrips(loggedInPassenger.passengerID)
					case "9":
						break loggedIn
					}
				}
			} else {
				fmt.Println("No such account")
			}
		}
	}

}

// Updating of Passenger Account
func UpdatePassengerAccount() {
	var newPassenger Passengers

	fmt.Print("Enter New First Name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newPassenger.firstName = strings.TrimSpace(input)

	fmt.Print("Enter New Last Name: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newPassenger.lastName = strings.TrimSpace(input)

	fmt.Print("Enter New Mobile Number: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	x := strings.TrimSpace(input)
	newPassenger.mobileNo, _ = strconv.Atoi(x)

	fmt.Print("Enter New Email Address: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newPassenger.emailAdd = strings.TrimSpace(input)

	client := &http.Client{}
	url := "https://localhost:5000/Passenger"
	postBody, _ := json.Marshal(newPassenger)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest("PATCH", url, resBody); err == nil {
		if res, err2 := client.Do(req); err2 == nil {
			if res.StatusCode == 202 {
				fmt.Println("Passenger account has been successfully modified")
			} else if res.StatusCode == 400 {
				fmt.Println("Bad Request")
			}
		}
	}
	// //Calling of database
	// db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

	// // Error handling
	// if err != nil {
	// 	fmt.Println("Error in connecting to database")
	// 	panic(err.Error())
	// }
	// defer db.Close()

	// //Inserting values into database
	// Stmt, err := db.Prepare("update passengers set firstName = ?,lastName = ?, mobileNo = ?, emailAdd = ? where passengerId = ?")

	// if err != nil {
	// 	fmt.Println("Error with sending data to database")
	// 	panic(err.Error())

	// }
	// defer Stmt.Close()
	// _, err = Stmt.Exec(newPassenger.firstName, newPassenger.lastName, newPassenger.mobileNo, newPassenger.emailAdd, loggedInPassenger.passengerID)
	// if err != nil {
	// 	fmt.Println("Error with sending data to database")
	// 	panic(err.Error())

	// } else {
	// 	// To notify of successful account creation
	// 	fmt.Println("====================")
	// 	fmt.Println("Passenger account has been successfully modified")
	// }
}
