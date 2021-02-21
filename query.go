package main

import (
	"database/sql"
)

func GetDonors(db *sql.DB) ([]Donor, error) {
	rows, err := db.Query(
		"SELECT donor_donor.id, donor_donor.phone, donor_donor.voter_id, donor_donor.dob, donor_donor.blood_group, donor_donor.disease, donor_donor.location, auth_user.first_name, auth_user.username, auth_user.email FROM PUBLIC.donor_donor INNER JOIN auth_user  on auth_user.id = donor_donor.donor_data_id")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var donors []Donor

	for rows.Next() {
		var donor Donor
		if err := rows.Scan(&donor.Id, &donor.Phone, &donor.VoterID, &donor.Dob, &donor.Blood, &donor.Disease, &donor.Location, &donor.FirstName, &donor.UserName, &donor.Email); err != nil {
			return nil, err
		}
		donors = append(donors, donor)
	}

	return donors, nil
}

func (donor *Donor) GetDonor(db *sql.DB) error {
	return db.QueryRow("SELECT donor_donor.id, donor_donor.phone, donor_donor.voter_id, donor_donor.dob, donor_donor.blood_group, donor_donor.disease, donor_donor.location, auth_user.first_name, auth_user.username, auth_user.email FROM PUBLIC.donor_donor INNER JOIN auth_user  on auth_user.id = donor_donor.donor_data_id WHERE donor_donor.id=$1",
		donor.Id).Scan(&donor.Id, &donor.Phone, &donor.VoterID, &donor.Dob, &donor.Blood, &donor.Disease, &donor.Location, &donor.FirstName, &donor.UserName, &donor.Email)
}

/*

func (donor *Donor) CreateDonor(db *sql.DB) error {
	err := db.QueryRow(
	"INSERT INTO PUBLIC.donor_donor(voter_id, phone, dob, blood_group, disease, location, donor_register_date, donor_data_id) VALUES($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
	 &donor.VoterID, &donor.Phone, &donor.Dob, &donor.Blood, &donor.Disease, &donor.Location, time.Now(), &donor.UserID).Scan(&donor.Id)

	_, err = db.Exec("UPDATE PUBLIC.donor_userprofile SET role='Donor' WHERE user_data_id=$1", &donor.UserID)

	if err != nil {
		return err
	}

	return nil
}

func (donor *Donor) UpdateDonor(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE PUBLIC.donor_donor SET phone=$1, disease=$2, location=$3  WHERE id=$4",
			&donor.Phone, &donor.Disease, &donor.Location, donor.Id)

	return err
}

func (donor *Donor) DeleteDonor(db *sql.DB) error {

	var temp = db.QueryRow("SELECT donor_data_id FROM PUBLIC.donor_donor WHERE id=$1", donor.Id)
	_, err := db.Exec("UPDATE PUBLIC.donor_userprofile SET role='Patient' WHERE user_data_id=$1", temp)

	return err
} */
