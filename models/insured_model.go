package models

type Insured struct {
	ID            int    `json:"id"`
	ContactID     int    `json:"contact_id"`
	PaymentID     int    `json:"payment_id"`
	FullName      string `json:"full_name"`
	DateOfBirth   string `json:"date_of_birth"`
	Address       string `json:"address"`
	StreetAddress string `json:"street_address"`
	City          string `json:"city"`
	Country       string `json:"country"`
	State         string `json:"state"`
	PostalCode    string `json:"postal_code"`
}
