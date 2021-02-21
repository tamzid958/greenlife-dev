package main

import "time"

type Donor struct {
	Id        int       `json:"id"`
	VoterID   string    `json:"voter_id"`
	Phone     string    `json:"phone"`
	Dob       string    `json:"dob"`
	Blood     string    `json:"blood"`
	Disease   bool      `json:"disease"`
	Location  string    `json:"location"`
	DonorReg  time.Time `json:"registration_date"`
	FirstName string    `json:"first_name"`
	UserName  string    `json:"username"`
	Email     string    `json:"email"`
}
