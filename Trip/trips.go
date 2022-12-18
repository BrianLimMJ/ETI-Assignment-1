package Trips

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type Trips struct {
	tripID      int
	passengerId string
	driverId    string
	startPostal string
	endPostal   string
	isStarted   bool
	isCompleted bool
	timeCreated time.Time
}

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

// Allows passenger's to request a new trip and assign driver who is has not taken up a job
func RequestTrip(passengerID int) {
	var newTrip Trips

	fmt.Print("Enter Starting Postal Code: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	newTrip.startPostal = strings.TrimSpace(input)

	fmt.Print("Enter Ending Postal Code: ")
	reader = bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	newTrip.endPostal = strings.TrimSpace(input)

	newTrip.timeCreated = time.Now()

	//Calling of database
	db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

	// Error handling
	if err != nil {
		fmt.Println("Error in connecting to database")
		panic(err.Error())
	}
	defer db.Close()

	//Checking for value in database
	result, err := db.Query("select * from drivers where isDriving = 0")
	if err != nil {
		fmt.Println("Error in searching for data")
		panic(err.Error())

		// To allow the program to assign a driver who is not driving to a trip
	} else {
		for result.Next() {
			var DriverInfo Drivers
			err = result.Scan(&DriverInfo.driverId, &DriverInfo.firstName, &DriverInfo.lastName, &DriverInfo.mobileNo, &DriverInfo.emailAdd, &DriverInfo.identificationNo, &DriverInfo.licenseNo, &DriverInfo.isDriving)
			if err != nil {
				panic(err.Error())

			} else {
				_, err = db.Exec("insert into trips (passengerID, driverID, startPostal, endPostal, timeCreated) values(?, ?, ?, ?, ?)",
					passengerID, DriverInfo.driverId, newTrip.startPostal, newTrip.endPostal, newTrip.timeCreated)
				if err != nil {
					fmt.Println("Error in sending data to database")
					panic(err.Error())

				} else {
					// To notify of successful account creation

					//Inserting values into database
					Stmt, err := db.Prepare("update drivers set isDriving = 1 where driverId = ?")

					if err != nil {
						fmt.Println("Error with sending data to database")
						panic(err.Error())

					}
					defer Stmt.Close()
					_, err = Stmt.Exec(DriverInfo.driverId)
					if err != nil {
						fmt.Println("Error with sending data to database")
						panic(err.Error())
					} else {
						fmt.Println("Driver is now booked")
						fmt.Println("The driver you have been assigned to is: ", DriverInfo.firstName)
					}

					fmt.Println("====================")
					fmt.Println("Trip has been successfully created")
				}
			}
		}
	}
}

// Checks available trips assigned to drivers using driverId
func CheckTrip(driverId int) {
	db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db?parseTime=true")

	// Error handling
	if err != nil {
		fmt.Println("Error in connecting to database")
		panic(err.Error())
	}
	defer db.Close()

	//Checking for value in database
	result, err := db.Query("select * from trips where driverId = ?", driverId)
	if err != nil {
		fmt.Println("Error with getting data from database")
		panic(err.Error())

	} else {
		for result.Next() {
			var newTrip Trips
			err = result.Scan(&newTrip.driverId, &newTrip.passengerId, &newTrip.driverId, &newTrip.startPostal, &newTrip.endPostal, &newTrip.isStarted, &newTrip.isCompleted, &newTrip.timeCreated)
			if err != nil {
				fmt.Println("No jobs available")
				panic(err.Error())

			} else {

				fmt.Println("====================")
				fmt.Println("Your starting and ending postal ", newTrip.startPostal, newTrip.endPostal)
			}
		}
	}
}

// Start's trip assigned to driver
func StartTrip(driverId int) {
	fmt.Print("Do you want to start trip? (Y/N): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	x := strings.ToLower(input)
	x = strings.TrimSpace(x)

	if x == "y" {
		//Calling of database
		db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

		// Error handling
		if err != nil {
			fmt.Println("Error in connecting to database")
			panic(err.Error())
		}
		defer db.Close()

		Stmt, err := db.Prepare("update trips set isStarted = 1, isCompleted = 0 where driverId = ?")

		if err != nil {
			fmt.Println("Error with sending data to database")
			panic(err.Error())

		}
		defer Stmt.Close()
		_, err = Stmt.Exec(driverId)
		if err != nil {
			fmt.Println("Error with sending data to database")
			panic(err.Error())

		} else {
			// To notify of successful account creation
			fmt.Println("====================")
			fmt.Println("You have started the trip")
		}

	} else if x == "n" {
		fmt.Println("You have selected no")
	} else {
		fmt.Println("Invalid input")
		fmt.Println("Exiting...")
	}
}

// Ends trip assigned to driver
func EndTrip(driverId int) {
	var newTrip Trips
	fmt.Print("Do you want to start trip? (Y/N): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	x := strings.ToLower(input)
	x = strings.TrimSpace(x)

	if x == "y" {
		//Calling of database
		db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

		// Error handling
		if err != nil {
			fmt.Println("Error in connecting to database")
			panic(err.Error())
		}
		defer db.Close()

		Stmt, err := db.Prepare("update trips set isStarted = 0, isCompleted = 1  where driverId = ?")

		if err != nil {
			fmt.Println("Error with sending data to database")
			panic(err.Error())

		}
		defer Stmt.Close()
		_, err = Stmt.Exec(driverId)
		if err != nil {
			fmt.Println("Error with sending data to database")
			panic(err.Error())

		} else if newTrip.isStarted != true {
			fmt.Println("====================")
			fmt.Println("You have ended the trip")
		} else {
			fmt.Println("There is nothing to end")

		}

		// To change the driver's availability into free to request trip
		Stmt, err = db.Prepare("update drivers set isDriving = 0 where driverId = ?")

		if err != nil {
			fmt.Println("Error with sending data to database")
			panic(err.Error())

		}
		defer Stmt.Close()
		_, err = Stmt.Exec(driverId)
		if err != nil {
			fmt.Println("Error with sending data to database")
			panic(err.Error())
		} else {
			fmt.Println("Thank you")
		}
	}
}

// Display all previous trips associated with Passenger and the passengerId
func RetriveTrips(passengerID int) {
	db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db?parseTime=true")

	// Error handling
	if err != nil {
		fmt.Println("Error in connecting to database")
		panic(err.Error())
	}
	defer db.Close()

	//Checking for value in database
	result, err := db.Query("select * from trips where passengerId = ? Order by passengerId Desc", passengerID)
	if err != nil {
		fmt.Println("Error with getting data from database")
		panic(err.Error())

	} else {
		for result.Next() {
			var newTrip Trips
			err = result.Scan(&newTrip.driverId, &newTrip.passengerId, &newTrip.driverId, &newTrip.startPostal, &newTrip.endPostal, &newTrip.isStarted, &newTrip.isCompleted, &newTrip.timeCreated)
			if err != nil {
				fmt.Println("No jobs available")
				panic(err.Error())

			} else {
				fmt.Println(passengerID, newTrip.startPostal, newTrip.endPostal, newTrip.timeCreated)
				continue
			}
		}
	}
}
