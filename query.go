package main

import "database/sql"

func GetDonors(db *sql.DB) ([]Donor, error) {
	rows, err := db.Query(
		"SELECT id, phone, dob, blood_group, disease, location FROM PUBLIC.donor_donor")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var donors []Donor

	for rows.Next() {
		var d Donor
		if err := rows.Scan(&d.Id, &d.Phone, &d.Dob, &d.Blood, &d.Disease, &d.Location); err != nil {
			return nil, err
		}
		donors = append(donors, d)
	}

	return donors, nil
}

func (donor *Donor) GetDonor(db *sql.DB) error {
	return db.QueryRow("SELECT id, phone, dob, blood_group, disease, location FROM PUBLIC.donor_donor WHERE id=$1",
		donor.Id).Scan(&donor.Id, &donor.Phone, &donor.Dob, &donor.Blood, &donor.Disease, &donor.Location)
}

func (donor *Donor) CreateDonor(db *sql.DB) error {
	err := db.QueryRow(
		"INSERT INTO PUBLIC.donor_donor(phone, dob, blood_group, disease, location) VALUES($1, $2, $3, $4, $5) RETURNING id",
		&donor.Phone, &donor.Dob, &donor.Blood, &donor.Disease, &donor.Location).Scan(&donor.Id)

	if err != nil {
		return err
	}

	return nil
}

func (donor *Donor) UpdateDonor(db *sql.DB) error {
	_, err :=
		db.Exec("UPDATE PUBLIC.donor_donor SET phone=$1, dob=$2, blood_group=$3, disease=$4, location=$5  WHERE id=$6",
			&donor.Phone, &donor.Dob, &donor.Blood, &donor.Disease, &donor.Location, donor.Id)

	return err
}

func (donor *Donor) DeleteDonor(db *sql.DB) error {
	_, err := db.Exec("DELETE FROM PUBLIC.donor_donor WHERE id=$1", donor.Id)

	return err
}
