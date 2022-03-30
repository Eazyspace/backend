package model

type Organization struct {
	OrgID           int64  `json:"orgId" gorm:"primaryKey"`
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	PhoneNumber     string `json:"phoneNumber,omitempty"`
	ImportancePoint int    `json:"importancePoint,omitempty"`
}
