package model

import (
	"gorm.io/gorm"
)

type Organization struct {
	gorm.Model
	OrganizationCode string `json:"organizationCode"`
	Name             string `json:"name,omitempty"`
	Email            string `json:"email,omitempty"`
	PhoneNumber      string `json:"phoneNumber,omitempty"`
	ImportancePoint  int    `json:"importancePoint,omitempty"`
}
