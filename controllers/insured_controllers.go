package controllers

import (
	"database/sql"
	"errors"
	"log"

	"github.com/vikash-parashar/01_billing/config"
	"github.com/vikash-parashar/01_billing/models"
)

func GetAllInsured() ([]models.Insured, error) {
	var insuredList []models.Insured

	rows, err := config.DB.Query("SELECT * FROM insured")
	if err != nil {
		log.Println("Error querying insured:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var insured models.Insured
		err := rows.Scan(&insured.ID, &insured.ContactID, &insured.FullName, &insured.DateOfBirth, &insured.Address, &insured.City, &insured.Country, &insured.State, &insured.ZipCode)
		if err != nil {
			log.Println("Error scanning insured rows:", err)
			return nil, err
		}
		insuredList = append(insuredList, insured)
	}

	return insuredList, nil
}

func CreateInsured(insured models.Insured) (int, error) {
	var id int
	err := config.DB.QueryRow("INSERT INTO insured (contact_id, full_name, date_of_birth, address, street_address, city, country, state, postal_code) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id",
		insured.ContactID, insured.FullName, insured.DateOfBirth, insured.Address, insured.City, insured.Country, insured.State, insured.ZipCode).Scan(&id)
	if err != nil {
		log.Println("Error creating insured:", err)
		return 0, err
	}

	return id, nil
}

func GetInsuredByID(id int) (models.Insured, error) {
	var insured models.Insured
	err := config.DB.QueryRow("SELECT * FROM insured WHERE id = $1", id).
		Scan(&insured.ID, &insured.ContactID, &insured.FullName, &insured.DateOfBirth, &insured.Address, &insured.City, &insured.Country, &insured.State, &insured.ZipCode)
	if err != nil {
		log.Println("Error fetching insured by ID:", err)
		if errors.Is(err, sql.ErrNoRows) {
			return models.Insured{}, errors.New("insured not found")
		}
		return models.Insured{}, err
	}

	return insured, nil
}

func UpdateInsuredByID(id int, insured models.Insured) error {
	result, err := config.DB.Exec("UPDATE insured SET contact_id=$1, full_name=$2, date_of_birth=$3, address=$4, street_address=$5, city=$6, country=$7, state=$8, postal_code=$9 WHERE id=$10",
		insured.ContactID, insured.FullName, insured.DateOfBirth, insured.Address, insured.City, insured.Country, insured.State, insured.ZipCode, id)
	if err != nil {
		log.Println("Error updating insured by ID:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("insured not found")
	}

	return nil
}

func DeleteInsuredByID(id int) error {
	result, err := config.DB.Exec("DELETE FROM insured WHERE id = $1", id)
	if err != nil {
		log.Println("Error deleting insured by ID:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("insured not found")
	}

	return nil
}
