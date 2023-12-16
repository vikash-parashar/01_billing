// controllers.go
package controllers

import (
	"database/sql"
	"time"

	_ "github.com/lib/pq"
	"github.com/vikash-parashar/01_billing/models"
)

var db *sql.DB // Initialize your PostgreSQL database connection

func GetAllContacts() ([]models.Contact, error) {
	var contacts []models.Contact
	rows, err := db.Query("SELECT * FROM contacts")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var contact models.Contact
		err := rows.Scan(
			&contact.ID,
			&contact.ContactType,
			&contact.EmailAddress,
			&contact.FullName,
			&contact.LocationID,
			&contact.OwnerUserID,
			&contact.PhoneNumber,
			&contact.PipelineID,
			&contact.PipelineStageID,
			&contact.CreatedAt,
			&contact.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		contacts = append(contacts, contact)
	}

	return contacts, nil
}

func CreateContact(contact models.Contact) (models.Contact, error) {
	query := `
		INSERT INTO contacts 
		(contactType, emailAddress, fullName, locationID, ownerUserID, phoneNumber, pipelineID, pipelineStageID, created_at, updated_at) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING id`

	err := db.QueryRow(
		query,
		contact.ContactType,
		contact.EmailAddress,
		contact.FullName,
		contact.LocationID,
		contact.OwnerUserID,
		contact.PhoneNumber,
		contact.PipelineID,
		contact.PipelineStageID,
		time.Now(), // created_at
		time.Now(), // updated_at
	).Scan(&contact.ID)

	if err != nil {
		return models.Contact{}, err
	}

	return contact, nil
}

func GetContactByID(contactID int) (models.Contact, error) {
	var contact models.Contact
	err := db.QueryRow("SELECT * FROM contacts WHERE id = $1", contactID).Scan(
		&contact.ID,
		&contact.ContactType,
		&contact.EmailAddress,
		&contact.FullName,
		&contact.LocationID,
		&contact.OwnerUserID,
		&contact.PhoneNumber,
		&contact.PipelineID,
		&contact.PipelineStageID,
		&contact.CreatedAt,
		&contact.UpdatedAt,
	)

	if err != nil {
		return models.Contact{}, err
	}

	return contact, nil
}

func UpdateContactByID(contactID int, updatedContact models.Contact) (models.Contact, error) {
	query := `
		UPDATE contacts 
		SET 
			contactType=$1, 
			emailAddress=$2, 
			fullName=$3, 
			locationID=$4, 
			ownerUserID=$5, 
			phoneNumber=$6, 
			pipelineID=$7, 
			pipelineStageID=$8, 
			updated_at=$9
		WHERE 
			id=$10
		RETURNING id`

	err := db.QueryRow(
		query,
		updatedContact.ContactType,
		updatedContact.EmailAddress,
		updatedContact.FullName,
		updatedContact.LocationID,
		updatedContact.OwnerUserID,
		updatedContact.PhoneNumber,
		updatedContact.PipelineID,
		updatedContact.PipelineStageID,
		time.Now(), // updated_at
		contactID,
	).Scan(&updatedContact.ID)

	if err != nil {
		return models.Contact{}, err
	}

	return updatedContact, nil
}

func DeleteContactByID(contactID int) error {
	_, err := db.Exec("DELETE FROM contacts WHERE id = $1", contactID)
	return err
}
