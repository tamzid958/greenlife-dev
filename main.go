package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Donor struct {
	Id       string    `json:"id"`
	Phone    string    `json:"phone"`
	Dob      string    `json:"dob"`
	Blood    string    `json:"blood"`
	Disease  bool      `json:"disease"`
	Location string    `json:"location"`
	Rating   string    `json:"rating"`
	Details  *Details  `json:"details"`
	Donation *Donation `json:"donation"`
}

type Details struct {
	FirstName string `json:"first_name"`
	UserName  string `json:"username"`
	Email     string `json:"email"`
}

type Donation struct {
	Review          int    `json:"review"`
	Status          bool   `json:"status"`
	DonationDate    string `json:"donation_date"`
	AppointmentDate string `json:"appointment_date"`
}

var _connStr = "postgres://postgres:12345@localhost/greenlife?sslmode=disable"

var donors = []Donor{
	{Id: "1", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Location: "location", Rating: "4", Details: &Details{FirstName: "name", UserName: "user", Email: "Email"}},
	{Id: "2", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Location: "location", Rating: "4"},
	{Id: "3", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Location: "location", Rating: "4"},
	{Id: "4", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Location: "location", Rating: "4"},
	{Id: "5", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Location: "location", Rating: "4"},
	{Id: "6", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Location: "location", Rating: "4"},
	{Id: "7", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Location: "location", Rating: "4"},
	{Id: "8", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Location: "location", Rating: "4"},
}

func GetDonors(writer http.ResponseWriter, _ *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	db, err := sql.Open("postgres", _connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT id, dob, phone, blood_group, disease, location FROM PUBLIC.donor_donor ORDER BY id ASC")

	json.NewEncoder(writer).Encode(donors)

	defer rows.Close()
	defer db.Close()
}
func GetDonor(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for _, item := range donors {
		if item.Id == params["id"] {
			json.NewEncoder(writer).Encode(item)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Donor{})
}
func CreateDonor(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	var donor Donor
	_ = json.NewDecoder(request.Body).Decode(&donor)
	donor.Id = strconv.Itoa(rand.Intn(100000))
	donors = append(donors, donor)
	json.NewEncoder(writer).Encode(donor)
}

func UpdateDonor(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range donors {
		if item.Id == params["id"] {
			donors = append(donors[:index], donors[index+1:]...)
			var donor Donor
			_ = json.NewDecoder(request.Body).Decode(&donor)
			donor.Id = params["id"]
			donors = append(donors, donor)
			json.NewEncoder(writer).Encode(donor)
			return
		}
	}
	json.NewEncoder(writer).Encode(donors)
}
func DeleteDonor(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(request)
	for index, item := range donors {
		if item.Id == params["id"] {
			donors = append(donors[:index], donors[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(donors)
}

func main() {
	r := mux.NewRouter()

	donors = append(donors, Donor{Id: "9", Phone: "0199999", Dob: "18 Jun 1982", Blood: "AB+", Disease: true, Location: "location", Rating: "4"})

	r.HandleFunc("/dev/donors", GetDonors).Methods("GET")
	r.HandleFunc("/dev/donors", CreateDonor).Methods("POST")

	r.HandleFunc("/dev/donors/{id:[0-9]+}", GetDonor).Methods("GET")
	r.HandleFunc("/dev/donors/{id:[0-9]+}", UpdateDonor).Methods("PUT")
	r.HandleFunc("/dev/donors/{id:[0-9]+}", DeleteDonor).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))
}
