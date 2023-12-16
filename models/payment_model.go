package models

import (
	"time"
)

// Payment struct to represent the database model
type Payment struct {
	// @json:id
	ID int `json:"id"`

	// @json:contact_id
	ContactID int `json:"contact_id"`

	// @json:payment_type
	PaymentType string `json:"payment_type"`

	// @json:status
	Status string `json:"status"`

	// @json:payer
	Payer string `json:"payer"`

	// @json:billing_order
	BillingOrder int `json:"billing_order"`

	// @json:condition_related_to
	ConditionRelatedTo string `json:"condition_related_to"`

	// @json:billing_id
	BillingID string `json:"billing_id"`

	// @json:relationship_to_insured
	RelationshipToInsured string `json:"relationship_to_insured"`

	// @json:insured_full_name
	InsuredFullName string `json:"insured_full_name"`

	// @json:insured_id
	InsuredID string `json:"insured_id"`

	// @json:dob_of_insured
	DOBOfInsured string `json:"dob_of_insured"`

	// @json:address_of_insured
	AddressOfInsured string `json:"address_of_insured"`

	// @json:city_of_insured
	CityOfInsured string `json:"city_of_insured"`

	// @json:state_of_insured
	StateOfInsured string `json:"state_of_insured"`

	// @json:zip_code_of_insured
	ZipCodeOfInsured string `json:"zip_code_of_insured"`

	// @json:insurance_start_date
	InsuranceStartDate string `json:"insurance_start_date"`

	// @json:insurance_end_date
	InsuranceEndDate string `json:"insurance_end_date"`

	// @json:insurance_co_pay
	InsuranceCoPay float64 `json:"insurance_co_pay"`

	// @json:insurance_deductible
	InsuranceDeductible float64 `json:"insurance_deductible"`

	// @json:created_at
	CreatedAt time.Time `json:"created_at"`

	// @json:updated_at
	UpdatedAt time.Time `json:"updated_at"`
}
