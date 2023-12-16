package models

type Insured struct {
	ID          int    `json:"id"`
	ContactID   int    `json:"contact_id"`
	FullName    string `json:"full_name"`
	DateOfBirth string `json:"date_of_birth"`
	Address     string `json:"address"`
	City        string `json:"city"`
	State       string `json:"state"`
	ZipCode     string `json:"postal_code"`
}
