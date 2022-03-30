package model

import (
	"time"
)

type User struct {
	UserID      int64     `json:"userId"`
	OrgID       int64     `json:"orgId"`
	DOB         time.Time `json:"dob,omitempty"`
	Role        int64     `json:"role,omitempty"`
	Faculty     string    `json:"faculty,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	AcademicID  string    `json:"academicID,omitempty"`
	Active      string    `json:"active,omitempty"`
	Activated   bool      `json:"activated,omitempty"`
}
