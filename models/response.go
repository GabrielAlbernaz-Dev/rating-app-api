package models

import "time"

type Response struct {
	ID              int       `json:"id"`
	ComplaintID     int       `json:"complaint_id"`
	UserID          *int      `json:"user_id,omitempty"`
	CompanyID       *int      `json:"company_id,omitempty"`
	Content         string    `json:"content"`
	CompanyResponse bool      `json:"company_response"`
	CreatedAt       time.Time `json:"created_at"`
}
