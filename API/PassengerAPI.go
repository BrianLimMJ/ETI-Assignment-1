package API

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type Passenger struct {
	PassengerID int    `json:"passengerId"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	MobileNo    int    `json:"mobileNo"`
	EmailAdd    string `json:"emailAdd"`
}

type Driver struct {
	DriverID         int    `json:"driverId"`
	FirstName        string `json:"firstName"`
	LastName         string `json:"lastName"`
	MobileNo         int    `json:"mobileNo"`
	EmailAdd         string `json:"emailAdd"`
	IdentificationNo string `json:"identificationNo"`
	LicenseNo        string `json:"licenseNo"`
	IsDriving        bool   `json:"isDriving"`
}

type Trip struct {
	TripID      int    `json:"tripId"`
	DriverID    string `json:"driverId"`
	PassengerID string `json:"passengerId"`
	StartPostal string `json:"startPostal"`
	EndPostal   string `json:"endPostal"`
	IsStarted   bool   `json:"isStarted"`
	IsCompleted bool   `json:"isCompleted"`
	TimeCreated bool   `json:"timeCreated"`
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/Passenger", PassengerMenu).Methods("POST", "PATCH")
	router.HandleFunc("/Driver", DriverMenu).Methods("POST", "PATCH")
	router.HandleFunc("/ShowJob", CheckTrip).Methods("GET")
	router.HandleFunc("/Trip", RetriveTrips).Methods("GET")
	fmt.Println("Listening at port 5000")
	log.Fatal(http.ListenAndServe(": 5000", router))
}

func PassengerMenu(w http.ResponseWriter, r *http.Request) {
	// Creation of Passenger Account
	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var newPassenger Passenger

			if err := json.Unmarshal(body, &newPassenger); err == nil {
				//Calling of database
				db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

				// Error handling
				if err != nil {
					fmt.Println("Error in connecting to database")
					panic(err.Error())
				}
				defer db.Close()

				//Inserting values into database
				_, err = db.Exec("insert into passengers (firstName, lastName, mobileNo, emailAdd) values(?, ?, ?, ?)",
					newPassenger.FirstName, newPassenger.LastName, newPassenger.MobileNo, newPassenger.EmailAdd)
				if err != nil {
					fmt.Println("Error with sending data to database")
					panic(err.Error())

				} else {
					// To notify of successful account creation
					fmt.Println("====================")
					fmt.Println("Passenger account has been successfully created")
				}
			}
		}
	} else if r.Method == "PATCH" {
		//Updating the passenger Account
		var UpdatePassenger Passenger
		err := json.NewDecoder(r.Body).Decode(&UpdatePassenger)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

		// Error handling
		if err != nil {
			fmt.Println("Error in connecting to database")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer db.Close()

		Stmt, err := db.Prepare("update passengers set firstName = ?,lastName = ?, mobileNo = ?, emailAdd = ? where passengerId = ?")

		if err != nil {
			fmt.Println("Error with sending data to database")
			http.Error(w, err.Error(), http.StatusBadRequest)

		}
		defer Stmt.Close()
		_, err = Stmt.Exec(UpdatePassenger.FirstName, UpdatePassenger.LastName, UpdatePassenger.MobileNo, UpdatePassenger.EmailAdd, UpdatePassenger.PassengerID)
		if err != nil {
			fmt.Println("Error with sending data to database")
			http.Error(w, err.Error(), http.StatusBadRequest)

		} else {
			// To notify of successful account creation
			fmt.Println("====================")
			fmt.Println("Passenger account has been successfully modified")
		}
	} else {
		http.Error(w, "Error: ", http.StatusBadRequest)
	}

}

func DriverMenu(w http.ResponseWriter, r *http.Request) {
	//Creating Driver Account
	if r.Method == "POST" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var newDriver Driver

			if err := json.Unmarshal(body, &newDriver); err == nil {
				//Calling of database
				db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

				// Error handling
				if err != nil {
					fmt.Println("Error in connecting to database")
					panic(err.Error())
				}
				defer db.Close()

				//Inserting values into database
				_, err = db.Exec("insert into drivers (firstName, lastName, mobileNo, emailAdd,identificationNo, licenseNo) values(?, ?, ?, ?,?,?)",
					newDriver.FirstName, newDriver.LastName, newDriver.MobileNo, newDriver.EmailAdd, newDriver.IdentificationNo, newDriver.LicenseNo)
				if err != nil {
					fmt.Println("Error with sending data to database")
					panic(err.Error())

				} else {
					// To notify of successful account creation
					fmt.Println("====================")
					fmt.Println("Driver account has been successfully created")
				}
			}
		}
	} else if r.Method == "PATCH" {
		//Updating the Driver Account
		var UpdateDriver Driver
		err := json.NewDecoder(r.Body).Decode(&UpdateDriver)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db")

		// Error handling
		if err != nil {
			fmt.Println("Error in connecting to database")
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		defer db.Close()

		Stmt, err := db.Prepare("update drivers set firstName = ?,lastName = ?, mobileNo = ?, emailAdd = ? identificationNo = ?, licensoNo =? where driverId = ?")

		if err != nil {
			fmt.Println("Error with sending data to database")
			http.Error(w, err.Error(), http.StatusBadRequest)

		}
		defer Stmt.Close()
		_, err = Stmt.Exec(UpdateDriver.FirstName, UpdateDriver.LastName, UpdateDriver.MobileNo, UpdateDriver.EmailAdd, UpdateDriver.IdentificationNo, UpdateDriver.LicenseNo, UpdateDriver.DriverID)
		if err != nil {
			fmt.Println("Error with sending data to database")
			http.Error(w, err.Error(), http.StatusBadRequest)

		} else {
			// To notify of successful account creation
			fmt.Println("====================")
			fmt.Println("Driver account has been successfully modified")
		}
	} else {
		http.Error(w, "Error: ", http.StatusBadRequest)
	}
}

func CheckTrip(w http.ResponseWriter, r *http.Request) {
	// Show all Trips assigned to Driver
	if r.Method == "GET" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var newTrip Trip

			if err := json.Unmarshal(body, &newTrip); err == nil {
				//Calling of database
				db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db?parseTime=true")

				// Error handling
				if err != nil {
					fmt.Println("Error in connecting to database")
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				defer db.Close()

				//Checking for value in database
				result, err := db.Query("select * from trips where driverId = ?", newTrip.DriverID)
				if err != nil {
					fmt.Println("Error with getting data from database")
					http.Error(w, err.Error(), http.StatusBadRequest)

				} else {
					for result.Next() {
						var newTrip Trip
						err = result.Scan(&newTrip.TripID, &newTrip.PassengerID, &newTrip.DriverID, &newTrip.StartPostal, &newTrip.EndPostal, &newTrip.IsStarted, &newTrip.IsCompleted, &newTrip.TimeCreated)
						if err != nil {
							fmt.Println("No jobs available")
							http.Error(w, err.Error(), http.StatusBadRequest)

						} else {

							fmt.Println("====================")
							fmt.Println("Your starting and ending postal ", newTrip.StartPostal, newTrip.EndPostal)
						}
					}
				}
			}
		}
	}
}

func RetriveTrips(w http.ResponseWriter, r *http.Request) {
	// Show all trips done by passenger
	if r.Method == "GET" {
		if body, err := ioutil.ReadAll(r.Body); err == nil {
			var newTrip Trip

			if err := json.Unmarshal(body, &newTrip); err == nil {
				//Calling of database
				db, err := sql.Open("mysql", "root:Pa$$w0rd@tcp(127.0.0.1:3306)/my_db?parseTime=true")

				// Error handling
				if err != nil {
					fmt.Println("Error in connecting to database")
					http.Error(w, err.Error(), http.StatusBadRequest)
				}
				defer db.Close()

				//Checking for value in database
				result, err := db.Query("select * from trips where passengerId = ? Order by passengerId Desc", newTrip.PassengerID)
				if err != nil {
					fmt.Println("Error with getting data from database")
					http.Error(w, err.Error(), http.StatusBadRequest)

				} else {
					for result.Next() {
						var newTrip Trip
						err = result.Scan(&newTrip.TripID, &newTrip.PassengerID, &newTrip.DriverID, &newTrip.StartPostal, &newTrip.EndPostal, &newTrip.IsStarted, &newTrip.IsCompleted, &newTrip.TimeCreated)
						if err != nil {
							fmt.Println("No jobs available")
							http.Error(w, err.Error(), http.StatusBadRequest)

						} else {
							fmt.Println(newTrip.PassengerID, newTrip.StartPostal, newTrip.EndPostal, newTrip.TimeCreated)
							continue
						}
					}
				}
			}
		}
	}
}
