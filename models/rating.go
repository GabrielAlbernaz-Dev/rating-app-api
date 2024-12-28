package models

import "time"

type Rating struct {
	ID             int       `json:"id"`
	ComplaintID    int       `json:"complaint_id"`
	UserID         int       `json:"user_id"`
	RatingValue    int       `json:"rating_value"`
	Comment        string    `json:"comment,omitempty"`
	CreationDate   time.Time `json:"creation_date"`
}