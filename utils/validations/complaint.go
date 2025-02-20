package validations

import (
	"strings"

	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/utils"
)


func ValidateComplaintCreateBody(complaint models.Complaint) error {
	if complaint.UserID == 0 {
		return utils.ErrorField("user_id", "required")
	}

	if complaint.CompanyID == 0 {
		return utils.ErrorField("company_id", "required")
	}

	if strings.TrimSpace(complaint.Title) == "" {
		return utils.ErrorField("title", "required")
	}

	if strings.TrimSpace(complaint.Description) == "" {
		return utils.ErrorField("description", "required")
	}

	return nil
}

func ValidateComplaintUpdateBody(complaint models.Complaint) error {
	if complaint.UserID != 0 {
		return utils.ErrorField("user_id","read-only and cannot be updated")
	}

	if complaint.CompanyID != 0 {
		return utils.ErrorField("company_id","read-only and cannot be updated")
	}

	return nil
}