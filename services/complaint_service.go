package services

import (
	"fmt"

	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/repositories"
)

func GetAllComplaints() ([]models.Complaint, error) {
	complaints, err := repositories.FindAllComplaints()
	if err != nil {
		return nil, err
	}

	return complaints, nil
}

func GetComplaint(complaint models.Complaint) (*models.Complaint, error) {
	foundComplaint, err := repositories.FindComplaint(complaint)
	if err != nil {
		return nil, err
	}

	return foundComplaint, nil
}

func CreateComplaint(complaint models.Complaint) error {
	return repositories.CreateComplaint(complaint)
}

func UpdateComplaint(complaint models.Complaint) error {
	foundComplaint, _ := repositories.FindComplaint(complaint)

	if foundComplaint == nil || foundComplaint.ID == 0 {
		return fmt.Errorf("complaint with title %s not found", complaint.Title)
	}
	
	return repositories.UpdateComplaint(complaint)
}

func DeleteComplaint(id int) error {
	var complaint models.Complaint = models.Complaint{ID: id}

	foundComplaint, err := repositories.FindComplaint(complaint)
	if err != nil {
		return err
	}

	if foundComplaint.ID == 0 {
		return fmt.Errorf("complaint with title %s not found", foundComplaint.Title)
	}

	return repositories.DeleteComplaint(*foundComplaint)
}