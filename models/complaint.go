package models

import "time"

type Complaint struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	CompanyID      int       `json:"company_id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	CreationDate   time.Time `json:"creation_date"`
	Status         string    `json:"status"`
	Visibility     string    `json:"visibility"`
	Rating         *int      `json:"rating,omitempty"`
}