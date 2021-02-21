package main

type Donor struct {
	Id       int       `json:"id"`
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
