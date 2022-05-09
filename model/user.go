package model

import (
	"time"
)

type User struct {
	UserID         int64        `json:"userId" gorm:"primaryKey;autoIncrement"`
	Name           string       `json:"name,omitempty"`
	OrganizationID int64        `json:"organizationId,omitempty"`
	Organization   Organization `json:"-"`
	DOB            time.Time    `json:"dob,omitempty"`
	Role           int64        `json:"role,omitempty"`
	Faculty        string       `json:"faculty,omitempty"`
	Email          string       `json:"email,omitempty"`
	PhoneNumber    string       `json:"phoneNumber,omitempty"`
	AcademicID     string       `json:"academicId,omitempty" gorm:"unique"`
	Password       string       `json:"password,omitempty"`
	Avatar         string       `json:"avatar,omitempty"`
	IsActivated    bool         `json:"isActivated,omitempty"`
}
