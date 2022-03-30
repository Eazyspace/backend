package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserCode    string    `json:"userCode"`
	DOB         time.Time `json:"dob,omitempty"`
	Role        int64     `json:"role,omitempty"`
	Faculty     string    `json:"faculty,omitempty"`
	Email       string    `json:"email,omitempty"`
	PhoneNumber string    `json:"phoneNumber,omitempty"`
	AcademicID  string    `json:"academicID,omitempty"`
	Active      string    `json:"active,omitempty"`
	Activated   bool      `json:"activated,omitempty"`
}
