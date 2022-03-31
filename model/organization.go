package model

type Organization struct {
	OrganizationID  int64  `json:"organizationId" gorm:"primaryKey"`
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	ImportancePoint int    `json:"importancePoint,omitempty"`
}
