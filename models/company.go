package models

import "time"

type Company struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CNPJ       string    `json:"cnpj"`
	Address    string    `json:"address,omitempty"`
	Email      string    `json:"email,omitempty"`
	Phone      string    `json:"phone,omitempty"`
	CategoryID *int      `json:"category_id,omitempty"`
	Active     bool      `json:"active"`
	CreatedAt  time.Time `json:"created_at"`
}
