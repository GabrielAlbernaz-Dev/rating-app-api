package models

import "time"

type Complaint struct {
	ID          int       `json:"id"`
	UserID      int       `json:"user_id"`
	CompanyID   int       `json:"company_id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Active      bool      `json:"active"`
	CreatedAt   time.Time `json:"created_at"`
}
