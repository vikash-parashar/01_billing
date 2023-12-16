package models

// Contact represents a contact model
type Contact struct {
	ID              int    `json:"id"`
	ContactType     string `json:"contactType"`
	EmailAddress    string `json:"emailAddress"`
	FullName        string `json:"fullName"`
	LocationID      string `json:"locationID"`
	OwnerUserID     string `json:"ownerUserID"`
	PhoneNumber     string `json:"phoneNumber"`
	PipelineID      string `json:"pipelineID"`
	PipelineStageID string `json:"pipelineStageID"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
