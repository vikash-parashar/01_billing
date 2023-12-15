package models

// Payment struct to represent the database model
type Payment struct {
	ID                    int
	ContactID             int
	PaymentType           string
	Status                string
	Payer                 string
	BillingOrder          int
	ConditionRelatedTo    string
	BillingID             string
	RelationshipToInsured string
	InsuredFullName       string
	InsuredID             string
	DOBOfInsured          string
	AddressOfInsured      string
	CityOfInsured         string
	StateOfInsured        string
	ZipCodeOfInsured      string
	InsuranceStartDate    string
	InsuranceEndDate      string
	InsuranceCoPay        float64
	InsuranceDeductible   float64
}
