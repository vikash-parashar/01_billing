// controllers.go
package controllers

import (
	_ "github.com/lib/pq"
	"github.com/vikash-parashar/01_billing/models"
)

func GetAllInsureds() ([]models.Insured, error) {
	var insureds []models.Insured
	rows, err := db.Query("SELECT * FROM insured")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var insured models.Insured
		err := rows.Scan(
			&insured.ID,
			&insured.ContactID,
			&insured.FullName,
			&insured.DateOfBirth,
			&insured.Address,
			&insured.City,
			&insured.Country,
			&insured.State,
			&insured.ZipCode,
		)
		if err != nil {
			return nil, err
		}
		insureds = append(insureds, insured)
	}

	return insureds, nil
}

func CreateInsured(insured models.Insured) (models.Insured, error) {
	query := `
		INSERT INTO insureds 
		(contact_id, full_name, date_of_birth, address, city, country, state, postal_code) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err := db.QueryRow(
		query,
		insured.ContactID,
		insured.FullName,
		insured.DateOfBirth,
		insured.Address,
		insured.City,
		insured.Country,
		insured.State,
		insured.ZipCode,
	).Scan(&insured.ID)

	if err != nil {
		return models.Insured{}, err
	}

	return insured, nil
}

func GetInsuredByID(insuredID int) (models.Insured, error) {
	var insured models.Insured
	err := db.QueryRow("SELECT * FROM insured WHERE id = $1", insuredID).Scan(
		&insured.ID,
		&insured.ContactID,
		&insured.FullName,
		&insured.DateOfBirth,
		&insured.Address,
		&insured.City,
		&insured.Country,
		&insured.State,
		&insured.ZipCode,
	)

	if err != nil {
		return models.Insured{}, err
	}

	return insured, nil
}

func UpdateInsuredByID(insuredID int, updatedInsured models.Insured) (models.Insured, error) {
	query := `
		UPDATE insureds 
		SET 
			contact_id=$1, 
			full_name=$2, 
			date_of_birth=$3, 
			address=$4, 
			city=$5, 
			country=$6, 
			state=$7, 
			postal_code=$8
		WHERE 
			id=$9
		RETURNING id`

	err := db.QueryRow(
		query,
		updatedInsured.ContactID,
		updatedInsured.FullName,
		updatedInsured.DateOfBirth,
		updatedInsured.Address,
		updatedInsured.City,
		updatedInsured.Country,
		updatedInsured.State,
		updatedInsured.ZipCode,
		insuredID,
	).Scan(&updatedInsured.ID)

	if err != nil {
		return models.Insured{}, err
	}

	return updatedInsured, nil
}

func DeleteInsuredByID(insuredID int) error {
	_, err := db.Exec("DELETE FROM insured WHERE id = $1", insuredID)
	return err
}

func GetAllInsuredsByContactID(contactID int) ([]models.Insured, error) {
	var insureds []models.Insured
	rows, err := db.Query("SELECT * FROM insured WHERE contact_id = $1", contactID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var insured models.Insured
		err := rows.Scan(
			&insured.ID,
			&insured.ContactID,
			&insured.FullName,
			&insured.DateOfBirth,
			&insured.Address,
			&insured.City,
			&insured.Country,
			&insured.State,
			&insured.ZipCode,
		)
		if err != nil {
			return nil, err
		}
		insureds = append(insureds, insured)
	}

	return insureds, nil
}

func CreateInsuredByContactID(contactID int, insured models.Insured) (models.Insured, error) {
	query := `
		INSERT INTO insureds 
		(contact_id, full_name, date_of_birth, address, city, country, state, postal_code) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id`

	err := db.QueryRow(
		query,
		contactID,
		insured.FullName,
		insured.DateOfBirth,
		insured.Address,
		insured.City,
		insured.Country,
		insured.State,
		insured.ZipCode,
	).Scan(&insured.ID)

	if err != nil {
		return models.Insured{}, err
	}

	return insured, nil
}

func GetInsuredByContactAndInsuredID(contactID, insuredID int) (models.Insured, error) {
	var insured models.Insured
	err := db.QueryRow("SELECT * FROM insured WHERE contact_id = $1 AND id = $2", contactID, insuredID).Scan(
		&insured.ID,
		&insured.ContactID,
		&insured.FullName,
		&insured.DateOfBirth,
		&insured.Address,
		&insured.City,
		&insured.Country,
		&insured.State,
		&insured.ZipCode,
	)

	if err != nil {
		return models.Insured{}, err
	}

	return insured, nil
}

func UpdateInsuredByContactAndInsuredID(contactID, insuredID int, updatedInsured models.Insured) (models.Insured, error) {
	query := `
		UPDATE insureds 
		SET 
			full_name=$1, 
			date_of_birth=$2, 
			address=$3, 
			city=$4, 
			country=$5, 
			state=$6, 
			postal_code=$7
		WHERE 
			contact_id=$8 AND id=$9
		RETURNING id`

	err := db.QueryRow(
		query,
		updatedInsured.FullName,
		updatedInsured.DateOfBirth,
		updatedInsured.Address,
		updatedInsured.City,
		updatedInsured.Country,
		updatedInsured.State,
		updatedInsured.ZipCode,
		contactID,
		insuredID,
	).Scan(&updatedInsured.ID)

	if err != nil {
		return models.Insured{}, err
	}

	return updatedInsured, nil
}

func DeleteInsuredByContactAndInsuredID(contactID, insuredID int) error {
	_, err := db.Exec("DELETE FROM insured WHERE contact_id = $1 AND id = $2", contactID, insuredID)
	return err
}
