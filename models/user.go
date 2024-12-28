package models

import (
	"time"
)

type User struct {
	ID               int        `json:"id"`
	Name             string     `json:"name"`
	Email            string     `json:"email"`
	Password         string     `json:"password"`
	RegistrationDate time.Time  `json:"registration_date"`
	LastLogin        *time.Time `json:"last_login,omitempty"`
	Active           bool       `json:"active"`
}
