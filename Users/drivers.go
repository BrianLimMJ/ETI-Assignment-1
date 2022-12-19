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

type Drivers struct {
	driverId         int
	firstName        string
	lastName         string
	mobileNo         int
	emailAdd         string
	identificationNo string
	licenseNo        string
	isDriving        bool
}

var loggedInDriver Drivers

func DriverMenu() {
outer:
	for {
		fmt.Println("====================")
		fmt.Println("Driver Menu\n",
			"1.Create Driver Account\n",
			"2.Login to Driver Account\n",
			"9.Quit\n")
		fmt.Println("====================")

		fmt.Print("Enter an option: ")
		var choice string
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(input)

		switch choice {
		case "1":
			CreateDriverAccount()
		case "2":
			LogInAsDriver()
		case "9":
			break outer
		}
	}
}

// Creating new Driver account
func CreateDriverAccount() {
	var newDriver Drivers

	fmt.Print("Enter First Name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newDriver.firstName = strings.TrimSpace(input)

	fmt.Print("Enter Last Name: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newDriver.lastName = strings.TrimSpace(input)

	fmt.Print("Enter Mobile Number: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	x := strings.TrimSpace(input)
	newDriver.mobileNo, _ = strconv.Atoi(x)

	fmt.Print("Enter Email Address: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newDriver.emailAdd = strings.TrimSpace(input)

	fmt.Print("Enter Identification Number ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newDriver.identificationNo = strings.TrimSpace(input)

	fmt.Print("Enter License Number: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newDriver.licenseNo = strings.TrimSpace(input)

	client := &http.Client{}
	url := "https://localhost:5000/Driver"
	postBody, _ := json.Marshal(newDriver)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest("POST", url, resBody); err == nil {
		if res, err2 := client.Do(req); err2 == nil {
			if res.StatusCode == 202 {
				fmt.Println("Driver account has been successfully created")
			} else if res.StatusCode == 400 {
				fmt.Println("Bad Request")
			}
		}
	}

	// db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

	// // Error handling
	// if err != nil {
	// 	fmt.Println("Error in connecting to database")
	// 	panic(err.Error())
	// }
	// defer db.Close()

	// //Inserting data into database
	// _, err = db.Exec("insert into drivers (firstName, lastName, mobileNo, emailAdd, identificationNo, licenseNo) values(?, ?, ?, ?, ?, ?)",
	// 	newDriver.firstName, newDriver.lastName, newDriver.mobileNo, newDriver.emailAdd, newDriver.identificationNo, newDriver.licenseNo)
	// if err != nil {
	// 	fmt.Println("Error in sending data to database")
	// 	panic(err.Error())

	// } else {
	// 	// To notify of successful account creation
	// 	fmt.Println("====================")
	// 	fmt.Println("Driver account has been successfully created")
	// }

}

// Logging in as Driver
func LogInAsDriver() {
	fmt.Print("Enter your mobile number: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	DriverMobileNo := strings.TrimSpace(input)

	//Calling of database
	db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

	// Error handling
	if err != nil {
		fmt.Println("Error in connecting to database")
		panic(err.Error())
	}
	defer db.Close()

	//Checking for value in database
	result, err := db.Query("select * from drivers where mobileNo = ?", DriverMobileNo)
	if err != nil {
		fmt.Println("Error with getting data from database")
		panic(err.Error())

	} else {
		for result.Next() {
			var DriverInfo Drivers
			err = result.Scan(&DriverInfo.driverId, &DriverInfo.firstName, &DriverInfo.lastName, &DriverInfo.mobileNo, &DriverInfo.emailAdd, &DriverInfo.identificationNo, &DriverInfo.licenseNo, &DriverInfo.isDriving)
			x := strconv.Itoa(DriverInfo.mobileNo)
			if err != nil {
				fmt.Println("Unsuccessful Login")
				panic(err.Error())

			} else if x == DriverMobileNo {
				loggedInDriver = DriverInfo
				// To notify of successful login
				fmt.Println("====================")
				fmt.Println("Successfully logged into " + DriverInfo.firstName)
				// Logged in as Driver section
			loggedIn:
				for {
					fmt.Print("\n")
					fmt.Println("====================")
					fmt.Println("You are now logged in as a Driver\n",
						"1.Update Driver Account\n",
						"2.Check for assigned trips\n",
						"3.Start Trip\n",
						"4.End Trip\n",
						"9.Quit\n")
					fmt.Println("====================")

					fmt.Print("Enter an option: ")
					var choice string
					reader := bufio.NewReader(os.Stdin)
					input, _ := reader.ReadString('\n')
					choice = strings.TrimSpace(input)

					switch choice {
					case "1":
						UpdateDriverAccount()
					case "2":
						Trips.CheckTrip(loggedInDriver.driverId)
					case "3":
						Trips.StartTrip(loggedInDriver.driverId)
					case "4":
						Trips.EndTrip(loggedInDriver.driverId)
					case "9":
						break loggedIn
					}
				}
			} else {

			}
		}
	}

}

// Updating of Driver account
func UpdateDriverAccount() {
	var newDriver Drivers

	fmt.Print("Enter New First Name: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newDriver.firstName = strings.TrimSpace(input)

	fmt.Print("Enter New Last Name: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newDriver.lastName = strings.TrimSpace(input)

	fmt.Print("Enter New Mobile Number: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	x := strings.TrimSpace(input)
	newDriver.mobileNo, _ = strconv.Atoi(x)

	fmt.Print("Enter New Email Address: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newDriver.emailAdd = strings.TrimSpace(input)

	fmt.Print("Enter New License Number: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newDriver.licenseNo = strings.TrimSpace(input)

	client := &http.Client{}
	url := "https://localhost:5000/Driver"
	postBody, _ := json.Marshal(newDriver)
	resBody := bytes.NewBuffer(postBody)
	if req, err := http.NewRequest("PATCH", url, resBody); err == nil {
		if res, err2 := client.Do(req); err2 == nil {
			if res.StatusCode == 202 {
				fmt.Println("Driver account has been successfully modified")
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
	// Stmt, err := db.Prepare("update drivers set firstName = ?,lastName = ?, mobileNo = ?, emailAdd = ?, licenseNo = ? where driverId = ?")

	// if err != nil {
	// 	fmt.Println("Error with sending data to database")
	// 	panic(err.Error())

	// }
	// defer Stmt.Close()
	// _, err = Stmt.Exec(newDriver.firstName, newDriver.lastName, newDriver.mobileNo, newDriver.emailAdd, newDriver.licenseNo, loggedInDriver.driverId)
	// if err != nil {
	// 	fmt.Println("Error with sending data to database")
	// 	panic(err.Error())

	// } else {
	// 	// To notify of successful account creation
	// 	fmt.Println("====================")
	// 	fmt.Println("Driver account has been successfully modified")
	// }
}
