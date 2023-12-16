package controllers

import (
	"database/sql"
	"errors"
	"log"

	"github.com/vikash-parashar/01_billing/config"
	"github.com/vikash-parashar/01_billing/models"
)

// GetAllContacts retrieves all contacts from the database.
func GetAllContacts() ([]models.Contact, error) {
	var contacts []models.Contact

	rows, err := config.DB.Query("SELECT * FROM contacts")
	if err != nil {
		log.Println("Error querying contacts:", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var contact models.Contact
		err := rows.Scan(
			&contact.ID, &contact.ContactType, &contact.EmailAddress, &contact.FullName,
			&contact.LocationID, &contact.OwnerUserID, &contact.PhoneNumber,
			&contact.PipelineID, &contact.PipelineStageID, &contact.CreatedAt, &contact.UpdatedAt,
		)
		if err != nil {
			log.Println("Error scanning contact rows:", err)
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

// CreateContact creates a new contact in the database.
func CreateContact(contact models.Contact) (string, error) {
	var id string
	err := config.DB.QueryRow("INSERT INTO contacts (contactType, emailAddress, fullName, locationID, ownerUserID, phoneNumber, pipelineID, pipelineStageID) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id",
		contact.ContactType, contact.EmailAddress, contact.FullName, contact.LocationID,
		contact.OwnerUserID, contact.PhoneNumber, contact.PipelineID, contact.PipelineStageID).Scan(&id)
	if err != nil {
		log.Println("Error creating contact:", err)
		return "", err
	}

	return id, nil
}

// GetContactByID retrieves a specific contact by ID from the database.
func GetContactByID(id string) (models.Contact, error) {
	var contact models.Contact
	err := config.DB.QueryRow("SELECT * FROM contacts WHERE id = $1", id).
		Scan(&contact.ID, &contact.ContactType, &contact.EmailAddress, &contact.FullName,
			&contact.LocationID, &contact.OwnerUserID, &contact.PhoneNumber,
			&contact.PipelineID, &contact.PipelineStageID, &contact.CreatedAt, &contact.UpdatedAt)
	if err != nil {
		log.Println("Error fetching contact by ID:", err)
		if errors.Is(err, sql.ErrNoRows) {
			return models.Contact{}, errors.New("contact not found")
		}
		return models.Contact{}, err
	}

	return contact, nil
}

// UpdateContactByID updates a specific contact by ID in the database.
func UpdateContactByID(id string, updatedContact models.Contact) error {
	result, err := config.DB.Exec("UPDATE contacts SET contactType=$1, emailAddress=$2, fullName=$3, locationID=$4, ownerUserID=$5, phoneNumber=$6, pipelineID=$7, pipelineStageID=$8, updated_at=now() WHERE id=$9",
		updatedContact.ContactType, updatedContact.EmailAddress, updatedContact.FullName, updatedContact.LocationID,
		updatedContact.OwnerUserID, updatedContact.PhoneNumber, updatedContact.PipelineID, updatedContact.PipelineStageID, id)
	if err != nil {
		log.Println("Error updating contact by ID:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("contact not found")
	}

	return nil
}

// DeleteContactByID deletes a specific contact by ID from the database.
func DeleteContactByID(id string) error {
	result, err := config.DB.Exec("DELETE FROM contacts WHERE id = $1", id)
	if err != nil {
		log.Println("Error deleting contact by ID:", err)
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Println("Error getting rows affected:", err)
		return err
	}

	if rowsAffected == 0 {
		return errors.New("contact not found")
	}

	return nil
}
