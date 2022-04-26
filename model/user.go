package model

import (
	"time"
)

type User struct {
	UserID         int64 `json:"userId"`
	OrganizationID int64 `json:"organizationId"`
	Organization   Organization
	DOB            time.Time `json:"dob,omitempty"`
	Role           int64     `json:"role,omitempty"`
	Faculty        string    `json:"faculty,omitempty"`
	Email          string    `json:"email,omitempty"`
	PhoneNumber    string    `json:"phoneNumber,omitempty"`
	AcademicID     string    `json:"academicId,omitempty"`
	Password       string    `json:"password,omitempty"`
	IsActivated    bool      `json:"isActivated,omitempty"`
}
