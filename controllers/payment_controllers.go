// controllers.go
package controllers

import (
	"time"

	_ "github.com/lib/pq"
	"github.com/vikash-parashar/01_billing/models"
)

func GetAllPayments() ([]models.Payment, error) {
	var payments []models.Payment
	rows, err := db.Query("SELECT * FROM payments")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(
			&payment.ID,
			&payment.ContactID,
			&payment.PaymentType,
			&payment.Status,
			&payment.Payer,
			&payment.BillingOrder,
			&payment.ConditionRelatedTo,
			&payment.BillingID,
			&payment.RelationshipToInsured,
			&payment.InsuredFullName,
			&payment.InsuredID,
			&payment.DOBOfInsured,
			&payment.AddressOfInsured,
			&payment.CityOfInsured,
			&payment.StateOfInsured,
			&payment.ZipCodeOfInsured,
			&payment.InsuranceStartDate,
			&payment.InsuranceEndDate,
			&payment.InsuranceCoPay,
			&payment.InsuranceDeductible,
			&payment.CreatedAt,
			&payment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func CreatePayment(payment models.Payment) (models.Payment, error) {
	query := `
		INSERT INTO payments 
		(contact_id, payment_type, status, payer, billing_order, condition_related_to, billing_id, relationship_to_insured, insured_full_name, insured_id, dob_of_insured, address_of_insured, city_of_insured, state_of_insured, zip_code_of_insured, insurance_start_date, insurance_end_date, insurance_co_pay, insurance_deductible, created_at, updated_at) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
		RETURNING id`

	err := db.QueryRow(
		query,
		payment.ContactID,
		payment.PaymentType,
		payment.Status,
		payment.Payer,
		payment.BillingOrder,
		payment.ConditionRelatedTo,
		payment.BillingID,
		payment.RelationshipToInsured,
		payment.InsuredFullName,
		payment.InsuredID,
		payment.DOBOfInsured,
		payment.AddressOfInsured,
		payment.CityOfInsured,
		payment.StateOfInsured,
		payment.ZipCodeOfInsured,
		payment.InsuranceStartDate,
		payment.InsuranceEndDate,
		payment.InsuranceCoPay,
		payment.InsuranceDeductible,
		time.Now(), // created_at
		time.Now(), // updated_at
	).Scan(&payment.ID)

	if err != nil {
		return models.Payment{}, err
	}

	return payment, nil
}

func GetPaymentByID(paymentID int) (models.Payment, error) {
	var payment models.Payment
	err := db.QueryRow("SELECT * FROM payments WHERE id = $1", paymentID).Scan(
		&payment.ID,
		&payment.ContactID,
		&payment.PaymentType,
		&payment.Status,
		&payment.Payer,
		&payment.BillingOrder,
		&payment.ConditionRelatedTo,
		&payment.BillingID,
		&payment.RelationshipToInsured,
		&payment.InsuredFullName,
		&payment.InsuredID,
		&payment.DOBOfInsured,
		&payment.AddressOfInsured,
		&payment.CityOfInsured,
		&payment.StateOfInsured,
		&payment.ZipCodeOfInsured,
		&payment.InsuranceStartDate,
		&payment.InsuranceEndDate,
		&payment.InsuranceCoPay,
		&payment.InsuranceDeductible,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)

	if err != nil {
		return models.Payment{}, err
	}

	return payment, nil
}

func UpdatePaymentByID(paymentID int, updatedPayment models.Payment) (models.Payment, error) {
	query := `
		UPDATE payments 
		SET 
			contact_id=$1, 
			payment_type=$2, 
			status=$3, 
			payer=$4, 
			billing_order=$5, 
			condition_related_to=$6, 
			billing_id=$7, 
			relationship_to_insured=$8, 
			insured_full_name=$9, 
			insured_id=$10, 
			dob_of_insured=$11, 
			address_of_insured=$12, 
			city_of_insured=$13, 
			state_of_insured=$14, 
			zip_code_of_insured=$15, 
			insurance_start_date=$16, 
			insurance_end_date=$17, 
			insurance_co_pay=$18, 
			insurance_deductible=$19, 
			updated_at=$20
		WHERE 
			id=$21
		RETURNING id`

	err := db.QueryRow(
		query,
		updatedPayment.ContactID,
		updatedPayment.PaymentType,
		updatedPayment.Status,
		updatedPayment.Payer,
		updatedPayment.BillingOrder,
		updatedPayment.ConditionRelatedTo,
		updatedPayment.BillingID,
		updatedPayment.RelationshipToInsured,
		updatedPayment.InsuredFullName,
		updatedPayment.InsuredID,
		updatedPayment.DOBOfInsured,
		updatedPayment.AddressOfInsured,
		updatedPayment.CityOfInsured,
		updatedPayment.StateOfInsured,
		updatedPayment.ZipCodeOfInsured,
		updatedPayment.InsuranceStartDate,
		updatedPayment.InsuranceEndDate,
		updatedPayment.InsuranceCoPay,
		updatedPayment.InsuranceDeductible,
		time.Now(), // updated_at
		paymentID,
	).Scan(&updatedPayment.ID)

	if err != nil {
		return models.Payment{}, err
	}

	return updatedPayment, nil
}

func DeletePaymentByID(paymentID int) error {
	_, err := db.Exec("DELETE FROM payments WHERE id = $1", paymentID)
	return err
}

func GetAllPaymentsByContactID(contactID int) ([]models.Payment, error) {
	var payments []models.Payment
	rows, err := db.Query("SELECT * FROM payments WHERE contact_id = $1", contactID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(
			&payment.ID,
			&payment.ContactID,
			&payment.PaymentType,
			&payment.Status,
			&payment.Payer,
			&payment.BillingOrder,
			&payment.ConditionRelatedTo,
			&payment.BillingID,
			&payment.RelationshipToInsured,
			&payment.InsuredFullName,
			&payment.InsuredID,
			&payment.DOBOfInsured,
			&payment.AddressOfInsured,
			&payment.CityOfInsured,
			&payment.StateOfInsured,
			&payment.ZipCodeOfInsured,
			&payment.InsuranceStartDate,
			&payment.InsuranceEndDate,
			&payment.InsuranceCoPay,
			&payment.InsuranceDeductible,
			&payment.CreatedAt,
			&payment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func CreatePaymentByContactID(contactID int, payment models.Payment) (models.Payment, error) {
	query := `
		INSERT INTO payments 
		(contact_id, payment_type, status, payer, billing_order, condition_related_to, billing_id, relationship_to_insured, insured_full_name, insured_id, dob_of_insured, address_of_insured, city_of_insured, state_of_insured, zip_code_of_insured, insurance_start_date, insurance_end_date, insurance_co_pay, insurance_deductible, created_at, updated_at) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
		RETURNING id`

	err := db.QueryRow(
		query,
		contactID,
		payment.PaymentType,
		payment.Status,
		payment.Payer,
		payment.BillingOrder,
		payment.ConditionRelatedTo,
		payment.BillingID,
		payment.RelationshipToInsured,
		payment.InsuredFullName,
		payment.InsuredID,
		payment.DOBOfInsured,
		payment.AddressOfInsured,
		payment.CityOfInsured,
		payment.StateOfInsured,
		payment.ZipCodeOfInsured,
		payment.InsuranceStartDate,
		payment.InsuranceEndDate,
		payment.InsuranceCoPay,
		payment.InsuranceDeductible,
		time.Now(),
		time.Now(),
	).Scan(&payment.ID)

	if err != nil {
		return models.Payment{}, err
	}

	return payment, nil
}

func GetPaymentByContactAndPaymentID(contactID, paymentID int) (models.Payment, error) {
	var payment models.Payment
	err := db.QueryRow("SELECT * FROM payments WHERE contact_id = $1 AND id = $2", contactID, paymentID).Scan(
		&payment.ID,
		&payment.ContactID,
		&payment.PaymentType,
		&payment.Status,
		&payment.Payer,
		&payment.BillingOrder,
		&payment.ConditionRelatedTo,
		&payment.BillingID,
		&payment.RelationshipToInsured,
		&payment.InsuredFullName,
		&payment.InsuredID,
		&payment.DOBOfInsured,
		&payment.AddressOfInsured,
		&payment.CityOfInsured,
		&payment.StateOfInsured,
		&payment.ZipCodeOfInsured,
		&payment.InsuranceStartDate,
		&payment.InsuranceEndDate,
		&payment.InsuranceCoPay,
		&payment.InsuranceDeductible,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)

	if err != nil {
		return models.Payment{}, err
	}

	return payment, nil
}

func UpdatePaymentByContactAndPaymentID(contactID, paymentID int, updatedPayment models.Payment) (models.Payment, error) {
	query := `
		UPDATE payments 
		SET 
			payment_type=$1, 
			status=$2, 
			payer=$3, 
			billing_order=$4, 
			condition_related_to=$5, 
			billing_id=$6, 
			relationship_to_insured=$7, 
			insured_full_name=$8, 
			insured_id=$9, 
			dob_of_insured=$10, 
			address_of_insured=$11, 
			city_of_insured=$12, 
			state_of_insured=$13, 
			zip_code_of_insured=$14, 
			insurance_start_date=$15, 
			insurance_end_date=$16, 
			insurance_co_pay=$17, 
			insurance_deductible=$18, 
			updated_at=$19
		WHERE 
			contact_id=$20 AND id=$21
		RETURNING id`

	err := db.QueryRow(
		query,
		updatedPayment.PaymentType,
		updatedPayment.Status,
		updatedPayment.Payer,
		updatedPayment.BillingOrder,
		updatedPayment.ConditionRelatedTo,
		updatedPayment.BillingID,
		updatedPayment.RelationshipToInsured,
		updatedPayment.InsuredFullName,
		updatedPayment.InsuredID,
		updatedPayment.DOBOfInsured,
		updatedPayment.AddressOfInsured,
		updatedPayment.CityOfInsured,
		updatedPayment.StateOfInsured,
		updatedPayment.ZipCodeOfInsured,
		updatedPayment.InsuranceStartDate,
		updatedPayment.InsuranceEndDate,
		updatedPayment.InsuranceCoPay,
		updatedPayment.InsuranceDeductible,
		time.Now(), // updated_at
		contactID,
		paymentID,
	).Scan(&updatedPayment.ID)

	if err != nil {
		return models.Payment{}, err
	}

	return updatedPayment, nil
}

func DeletePaymentByContactAndPaymentID(contactID, paymentID int) error {
	_, err := db.Exec("DELETE FROM payments WHERE contact_id = $1 AND id = $2", contactID, paymentID)
	return err
}

func GetAllPaymentsByInsuredID(insuredID int) ([]models.Payment, error) {
	var payments []models.Payment
	rows, err := db.Query("SELECT * FROM payments WHERE insured_id = $1", insuredID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var payment models.Payment
		err := rows.Scan(
			&payment.ID,
			&payment.ContactID,
			&payment.PaymentType,
			&payment.Status,
			&payment.Payer,
			&payment.BillingOrder,
			&payment.ConditionRelatedTo,
			&payment.BillingID,
			&payment.RelationshipToInsured,
			&payment.InsuredFullName,
			&payment.InsuredID,
			&payment.DOBOfInsured,
			&payment.AddressOfInsured,
			&payment.CityOfInsured,
			&payment.StateOfInsured,
			&payment.ZipCodeOfInsured,
			&payment.InsuranceStartDate,
			&payment.InsuranceEndDate,
			&payment.InsuranceCoPay,
			&payment.InsuranceDeductible,
			&payment.CreatedAt,
			&payment.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		payments = append(payments, payment)
	}

	return payments, nil
}

func CreatePaymentByInsuredID(insuredID int, payment models.Payment) (models.Payment, error) {
	query := `
		INSERT INTO payments 
		(contact_id, payment_type, status, payer, billing_order, condition_related_to, billing_id, relationship_to_insured, insured_full_name, insured_id, dob_of_insured, address_of_insured, city_of_insured, state_of_insured, zip_code_of_insured, insurance_start_date, insurance_end_date, insurance_co_pay, insurance_deductible, created_at, updated_at) 
		VALUES 
		($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
		RETURNING id`

	err := db.QueryRow(
		query,
		payment.ContactID,
		payment.PaymentType,
		payment.Status,
		payment.Payer,
		payment.BillingOrder,
		payment.ConditionRelatedTo,
		payment.BillingID,
		payment.RelationshipToInsured,
		payment.InsuredFullName,
		payment.InsuredID,
		payment.DOBOfInsured,
		payment.AddressOfInsured,
		payment.CityOfInsured,
		payment.StateOfInsured,
		payment.ZipCodeOfInsured,
		payment.InsuranceStartDate,
		payment.InsuranceEndDate,
		payment.InsuranceCoPay,
		payment.InsuranceDeductible,
		time.Now(),
		time.Now(),
	).Scan(&payment.ID)

	if err != nil {
		return models.Payment{}, err
	}

	return payment, nil
}

func GetPaymentByPaymentAndInsuredID(insuredID, paymentID int) (models.Payment, error) {
	var payment models.Payment
	err := db.QueryRow("SELECT * FROM payments WHERE insured_id = $1 AND id = $2", insuredID, paymentID).Scan(
		&payment.ID,
		&payment.ContactID,
		&payment.PaymentType,
		&payment.Status,
		&payment.Payer,
		&payment.BillingOrder,
		&payment.ConditionRelatedTo,
		&payment.BillingID,
		&payment.RelationshipToInsured,
		&payment.InsuredFullName,
		&payment.InsuredID,
		&payment.DOBOfInsured,
		&payment.AddressOfInsured,
		&payment.CityOfInsured,
		&payment.StateOfInsured,
		&payment.ZipCodeOfInsured,
		&payment.InsuranceStartDate,
		&payment.InsuranceEndDate,
		&payment.InsuranceCoPay,
		&payment.InsuranceDeductible,
		&payment.CreatedAt,
		&payment.UpdatedAt,
	)

	if err != nil {
		return models.Payment{}, err
	}

	return payment, nil
}

func UpdatePaymentByPaymentAndInsuredID(insuredID, paymentID int, updatedPayment models.Payment) (models.Payment, error) {
	query := `
		UPDATE payments 
		SET 
			payment_type=$1, 
			status=$2, 
			payer=$3, 
			billing_order=$4, 
			condition_related_to=$5, 
			billing_id=$6, 
			relationship_to_insured=$7, 
			insured_full_name=$8, 
			insured_id=$9, 
			dob_of_insured=$10, 
			address_of_insured=$11, 
			city_of_insured=$12, 
			state_of_insured=$13, 
			zip_code_of_insured=$14, 
			insurance_start_date=$15, 
			insurance_end_date=$16, 
			insurance_co_pay=$17, 
			insurance_deductible=$18, 
			updated_at=$19
		WHERE 
			insured_id=$20 AND id=$21
		RETURNING id`

	err := db.QueryRow(
		query,
		updatedPayment.PaymentType,
		updatedPayment.Status,
		updatedPayment.Payer,
		updatedPayment.BillingOrder,
		updatedPayment.ConditionRelatedTo,
		updatedPayment.BillingID,
		updatedPayment.RelationshipToInsured,
		updatedPayment.InsuredFullName,
		updatedPayment.InsuredID,
		updatedPayment.DOBOfInsured,
		updatedPayment.AddressOfInsured,
		updatedPayment.CityOfInsured,
		updatedPayment.StateOfInsured,
		updatedPayment.ZipCodeOfInsured,
		updatedPayment.InsuranceStartDate,
		updatedPayment.InsuranceEndDate,
		updatedPayment.InsuranceCoPay,
		updatedPayment.InsuranceDeductible,
		time.Now(), // updated_at
		insuredID,
		paymentID,
	).Scan(&updatedPayment.ID)

	if err != nil {
		return models.Payment{}, err
	}

	return updatedPayment, nil
}

func DeletePaymentByPaymentAndInsuredID(insuredID, paymentID int) error {
	_, err := db.Exec("DELETE FROM payments WHERE insured_id = $1 AND id = $2", insuredID, paymentID)
	return err
}
