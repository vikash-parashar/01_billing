package controllers

import (
	"database/sql"
	"time"

	"github.com/vikash-parashar/01_billing/config"
	"github.com/vikash-parashar/01_billing/models"
)

// CreatePaymentMethod inserts a new payment method into the database.
func CreatePaymentMethod(contactID string, payment models.Payment) (sql.Result, error) {

	// Insert the payment method into the database
	res, err := config.DB.Exec(`
		INSERT INTO payments (
			contact_id, payment_type, status, payer, billing_order, 
			condition_related_to, billing_id, relationship_to_insured, 
			insured_full_name, insured_id, dob_of_insured, address_of_insured, 
			city_of_insured, state_of_insured, zip_code_of_insured, 
			insurance_start_date, insurance_end_date, insurance_co_pay, 
			insurance_deductible, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
	`,
		contactID, payment.PaymentType, payment.Status, payment.Payer, payment.BillingOrder,
		payment.ConditionRelatedTo, payment.BillingID, payment.RelationshipToInsured,
		payment.InsuredFullName, payment.InsuredID, payment.DOBOfInsured, payment.AddressOfInsured,
		payment.CityOfInsured, payment.StateOfInsured, payment.ZipCodeOfInsured,
		payment.InsuranceStartDate, payment.InsuranceEndDate, payment.InsuranceCoPay,
		payment.InsuranceDeductible, time.Now(), time.Now(),
	)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// GetAllPaymentMethodsForContact retrieves all payment methods for a given contact ID from the database.
func GetAllPaymentMethodsForContact(contactID string) ([]models.Payment, error) {

	// Query all payment methods for the given contact ID
	rows, err := config.DB.Query("SELECT * FROM payments WHERE contact_id = $1", contactID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Iterate through the result set and populate payment methods
	var paymentMethods []models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(
			&payment.ID, &payment.ContactID, &payment.PaymentType, &payment.Status, &payment.Payer,
			&payment.BillingOrder, &payment.ConditionRelatedTo, &payment.BillingID,
			&payment.RelationshipToInsured, &payment.InsuredFullName, &payment.InsuredID,
			&payment.DOBOfInsured, &payment.AddressOfInsured, &payment.CityOfInsured,
			&payment.StateOfInsured, &payment.ZipCodeOfInsured, &payment.InsuranceStartDate,
			&payment.InsuranceEndDate, &payment.InsuranceCoPay, &payment.InsuranceDeductible,
			&payment.CreatedAt, &payment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, payment)
	}

	return paymentMethods, nil
}

// GetPaymentMethodByMethodID retrieves a specific payment method for a given contact and payment method ID from the database.
func GetPaymentMethodByMethodID(contactID, methodID string) (models.Payment, error) {

	// Query the payment method for the given contact and payment method ID
	var paymentMethod models.Payment
	err := config.DB.QueryRow(`
		SELECT * FROM payments WHERE contact_id = $1 AND id = $2
	`, contactID, methodID).Scan(
		&paymentMethod.ID, &paymentMethod.ContactID, &paymentMethod.PaymentType, &paymentMethod.Status, &paymentMethod.Payer,
		&paymentMethod.BillingOrder, &paymentMethod.ConditionRelatedTo, &paymentMethod.BillingID,
		&paymentMethod.RelationshipToInsured, &paymentMethod.InsuredFullName, &paymentMethod.InsuredID,
		&paymentMethod.DOBOfInsured, &paymentMethod.AddressOfInsured, &paymentMethod.CityOfInsured,
		&paymentMethod.StateOfInsured, &paymentMethod.ZipCodeOfInsured, &paymentMethod.InsuranceStartDate,
		&paymentMethod.InsuranceEndDate, &paymentMethod.InsuranceCoPay, &paymentMethod.InsuranceDeductible,
		&paymentMethod.CreatedAt, &paymentMethod.UpdatedAt,
	)
	if err != nil {
		return models.Payment{}, err
	}

	return paymentMethod, nil
}

// DeletePaymentMethod deletes a specific payment method for a given contact and payment method ID from the database.
func DeletePaymentMethod(contactID, methodID string) error {

	// Delete the payment method for the given contact and payment method ID
	_, err := config.DB.Exec("DELETE FROM payments WHERE contact_id = $1 AND id = $2", contactID, methodID)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePaymentMethod updates a specific payment method for a given contact and payment method ID in the database.
func UpdatePaymentMethod(contactID, methodID string, updatedPayment models.Payment) error {

	// Fetch the original payment from the database
	originalPayment, err := GetPaymentMethodByMethodID(contactID, methodID)
	if err != nil {
		return err
	}

	// Update only non-nil or non-empty fields in the original payment with values from the updated payment
	if updatedPayment.PaymentType == "" {
		originalPayment.PaymentType = updatedPayment.PaymentType
	}
	if updatedPayment.Status == "" {
		originalPayment.Status = updatedPayment.Status
	}
	if updatedPayment.Payer == "" {
		originalPayment.Payer = updatedPayment.Payer
	}
	if updatedPayment.BillingOrder >= 0 {
		originalPayment.BillingOrder = updatedPayment.BillingOrder
	}
	if updatedPayment.ConditionRelatedTo == "" {
		originalPayment.ConditionRelatedTo = updatedPayment.ConditionRelatedTo
	}
	if updatedPayment.BillingID == "" {
		originalPayment.BillingID = updatedPayment.BillingID
	}
	if updatedPayment.RelationshipToInsured == "" {
		originalPayment.RelationshipToInsured = updatedPayment.RelationshipToInsured
	}
	if updatedPayment.InsuredFullName == "" {
		originalPayment.InsuredFullName = updatedPayment.InsuredFullName
	}
	if updatedPayment.InsuredID == "" {
		originalPayment.InsuredID = updatedPayment.InsuredID
	}
	if updatedPayment.DOBOfInsured == "" {
		originalPayment.DOBOfInsured = updatedPayment.RelationshipToInsured
	}
	if updatedPayment.AddressOfInsured == "" {
		originalPayment.AddressOfInsured = updatedPayment.AddressOfInsured
	}
	if updatedPayment.CityOfInsured == "" {
		originalPayment.CityOfInsured = updatedPayment.CityOfInsured
	}
	if updatedPayment.StateOfInsured == "" {
		originalPayment.StateOfInsured = updatedPayment.StateOfInsured
	}
	if updatedPayment.ZipCodeOfInsured == "" {
		originalPayment.ZipCodeOfInsured = updatedPayment.RelationshipToInsured
	}
	if updatedPayment.InsuranceStartDate == "" {
		originalPayment.InsuranceStartDate = updatedPayment.InsuranceStartDate
	}
	if updatedPayment.InsuranceEndDate == "" {
		originalPayment.InsuranceEndDate = updatedPayment.InsuranceEndDate
	}
	if updatedPayment.InsuranceCoPay == 0 {
		originalPayment.InsuranceCoPay = updatedPayment.InsuranceCoPay
	}
	if updatedPayment.InsuranceDeductible == 0 {
		originalPayment.InsuranceDeductible = updatedPayment.InsuranceDeductible
	}
	if updatedPayment.CreatedAt.String() == "" {
		originalPayment.CreatedAt = updatedPayment.CreatedAt
	}
	if updatedPayment.UpdatedAt.String() == "" {
		originalPayment.UpdatedAt = updatedPayment.UpdatedAt
	}

	// Prepare the UPDATE query
	query := `
		UPDATE payments SET
			payment_type = $1, status = $2, payer = $3, billing_order = $4,
			condition_related_to = $5, billing_id = $6, relationship_to_insured = $7,
			insured_full_name = $8, insured_id = $9, dob_of_insured = $10, 
			address_of_insured = $11, city_of_insured = $12, state_of_insured = $13,
			zip_code_of_insured = $14, insurance_start_date = $15, insurance_end_date = $16,
			insurance_co_pay = $17, insurance_deductible = $18, updated_at = $19
		WHERE contact_id = $20 AND id = $21
	`

	// Execute the UPDATE query
	_, err = config.DB.Exec(query,
		originalPayment.PaymentType, originalPayment.Status, originalPayment.Payer, originalPayment.BillingOrder,
		originalPayment.ConditionRelatedTo, originalPayment.BillingID, originalPayment.RelationshipToInsured,
		originalPayment.InsuredFullName, originalPayment.InsuredID, originalPayment.DOBOfInsured,
		originalPayment.AddressOfInsured, originalPayment.CityOfInsured, originalPayment.StateOfInsured,
		originalPayment.ZipCodeOfInsured, originalPayment.InsuranceStartDate, originalPayment.InsuranceEndDate,
		originalPayment.InsuranceCoPay, originalPayment.InsuranceDeductible, time.Now(), contactID, methodID,
	)
	if err != nil {
		return err
	}

	return nil
}

// CreatePaymentMethodForInsured inserts a new payment method for a given insured ID into the database.
func CreatePaymentMethodForInsured(insuredID string, payment models.Payment) (sql.Result, error) {
	res, err := config.DB.Exec(`
		INSERT INTO insured_payments (
			insured_id, payment_type, status, payer, billing_order, 
			condition_related_to, billing_id, relationship_to_contact, 
			contact_full_name, contact_id, dob_of_contact, address_of_contact, 
			city_of_contact, state_of_contact, zip_code_of_contact, 
			insurance_start_date, insurance_end_date, insurance_co_pay, 
			insurance_deductible, created_at, updated_at
		) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21)
	`,
		insuredID, payment.PaymentType, payment.Status, payment.Payer, payment.BillingOrder,
		payment.ConditionRelatedTo, payment.BillingID, payment.RelationshipToContact,
		payment.ContactFullName, payment.ContactID, payment.DOBOfContact, payment.AddressOfContact,
		payment.CityOfContact, payment.StateOfContact, payment.ZipCodeOfContact,
		payment.InsuranceStartDate, payment.InsuranceEndDate, payment.InsuranceCoPay,
		payment.InsuranceDeductible, time.Now(), time.Now(),
	)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAllPaymentMethodsForInsured retrieves all payment methods for a given insured ID from the database.
func GetAllPaymentMethodsForInsured(insuredID string) ([]models.Payment, error) {
	rows, err := config.DB.Query("SELECT * FROM insured_payments WHERE insured_id = $1", insuredID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var paymentMethods []models.Payment
	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(
			&payment.ID, &payment.InsuredID, &payment.PaymentType, &payment.Status, &payment.Payer,
			&payment.BillingOrder, &payment.ConditionRelatedTo, &payment.BillingID,
			&payment.RelationshipToContact, &payment.ContactFullName, &payment.ContactID,
			&payment.DOBOfContact, &payment.AddressOfContact, &payment.CityOfContact,
			&payment.StateOfContact, &payment.ZipCodeOfContact, &payment.InsuranceStartDate,
			&payment.InsuranceEndDate, &payment.InsuranceCoPay, &payment.InsuranceDeductible,
			&payment.CreatedAt, &payment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		paymentMethods = append(paymentMethods, payment)
	}
	return paymentMethods, nil
}

// GetPaymentMethodForInsuredByMethodID retrieves a specific payment method for a given insured and payment method ID from the database.
func GetPaymentMethodForInsuredByMethodID(insuredID, methodID string) (models.Payment, error) {
	var payment models.Payment
	err := config.DB.QueryRow(`
		SELECT * FROM insured_payments WHERE insured_id = $1 AND id = $2
	`, insuredID, methodID).Scan(
		&payment.ID, &payment.InsuredID, &payment.PaymentType, &payment.Status, &payment.Payer,
		&payment.BillingOrder, &payment.ConditionRelatedTo, &payment.BillingID,
		&payment.RelationshipToContact, &payment.ContactFullName, &payment.ContactID,
		&payment.DOBOfContact, &payment.AddressOfContact, &payment.CityOfContact,
		&payment.StateOfContact, &payment.ZipCodeOfContact, &payment.InsuranceStartDate,
		&payment.InsuranceEndDate, &payment.InsuranceCoPay, &payment.InsuranceDeductible,
		&payment.CreatedAt, &payment.UpdatedAt,
	)
	if err != nil {
		return models.Payment{}, err
	}
	return payment, nil
}

// DeletePaymentMethodForInsured deletes a specific payment method for a given insured and payment method ID from the database.
func DeletePaymentMethodForInsured(insuredID, methodID string) error {
	_, err := config.DB.Exec("DELETE FROM insured_payments WHERE insured_id = $1 AND id = $2", insuredID, methodID)
	if err != nil {
		return err
	}
	return nil
}

// UpdatePaymentMethodForInsured updates a specific payment method for a given insured and payment method ID in the database.
func UpdatePaymentMethodForInsured(insuredID, methodID string, updatedPayment models.Payment) error {
	originalPayment, err := GetPaymentMethodForInsuredByMethodID(insuredID, methodID)
	if err != nil {
		return err
	}

	if updatedPayment.PaymentType == "" {
		originalPayment.PaymentType = updatedPayment.PaymentType
	}
	// ... (update other fields similarly)

	query := `
		UPDATE insured_payments SET
			payment_type = $1, status = $2, payer = $3, billing_order = $4,
			condition_related_to = $5, billing_id = $6, relationship_to_contact = $7,
			contact_full_name = $8, contact_id = $9, dob_of_contact = $10, 
			address_of_contact = $11, city_of_contact = $12, state_of_contact = $13,
			zip_code_of_contact = $14, insurance_start_date = $15, insurance_end_date = $16,
			insurance_co_pay = $17, insurance_deductible = $18, updated_at = $19
		WHERE insured_id = $20 AND id = $21
	`

	_, err = config.DB.Exec(query,
		originalPayment.PaymentType, originalPayment.Status, originalPayment.Payer, originalPayment.BillingOrder,
		originalPayment.ConditionRelatedTo, originalPayment.BillingID, originalPayment.RelationshipToContact,
		originalPayment.ContactFullName, originalPayment.ContactID, originalPayment.DOBOfContact,
		originalPayment.AddressOfContact, originalPayment.CityOfContact, originalPayment.StateOfContact,
		originalPayment.ZipCodeOfContact, originalPayment.InsuranceStartDate, originalPayment.InsuranceEndDate,
		originalPayment.InsuranceCoPay, originalPayment.InsuranceDeductible, time.Now(), insuredID, methodID,
	)
	if err != nil {
		return err
	}
	return nil
}
