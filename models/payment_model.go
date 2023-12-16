package models

import (
	"time"
)

const (
	PaymentStatusActive   = "Active"
	PaymentStatusInactive = "Inactive"
)

const (
	PaymentTypeCash    = "Cash"
	PaymentTypeInsured = "Insurance"
)

// RelationshipToInsured represents the possible relationship options.
type RelationshipToInsuredType string

const (
	Self   RelationshipToInsuredType = "self"
	Spouse RelationshipToInsuredType = "spouse"
	Child  RelationshipToInsuredType = "child"
	Others RelationshipToInsuredType = "others"
)

type Payment struct {
	ID                    int       `json:"id"`
	ContactID             int       `json:"contact_id"`
	PaymentType           string    `json:"payment_type"`
	Status                string    `json:"status"`
	Payer                 string    `json:"payer"`
	BillingOrder          int       `json:"billing_order"`
	ConditionRelatedTo    string    `json:"condition_related_to"`
	BillingID             string    `json:"billing_id"`
	RelationshipToInsured string    `json:"relationship_to_insured"`
	InsuredFullName       string    `json:"insured_full_name"`
	InsuredID             string    `json:"insured_id"`
	DOBOfInsured          string    `json:"dob_of_insured"`
	AddressOfInsured      string    `json:"address_of_insured"`
	CityOfInsured         string    `json:"city_of_insured"`
	StateOfInsured        string    `json:"state_of_insured"`
	ZipCodeOfInsured      string    `json:"zip_code_of_insured"`
	InsuranceStartDate    string    `json:"insurance_start_date"`
	InsuranceEndDate      string    `json:"insurance_end_date"`
	InsuranceCoPay        float64   `json:"insurance_co_pay"`
	InsuranceDeductible   float64   `json:"insurance_deductible"`
	CreatedAt             time.Time `json:"created_at"`
	UpdatedAt             time.Time `json:"updated_at"`
}
