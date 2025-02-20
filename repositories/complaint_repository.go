package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gabrielalbernazdev/rating-app-api/infra/database"
	"github.com/gabrielalbernazdev/rating-app-api/models"
)

const (
	complaintTable        = "complaints"
	selectComplaintFields = "id, user_id, company_id, title, description, active, created_at"
	insertComplaintFields = "user_id, company_id, title, description, active, created_at"
)

func FindAllComplaints() ([]models.Complaint, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s", selectComplaintFields, complaintTable)

	rows, err := database.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("error to get complaints: %v", err)
	}
	defer rows.Close()

	var complaints []models.Complaint

	for rows.Next() {
		var complaint models.Complaint
		err := rows.Scan(
			&complaint.ID,
			&complaint.UserID,
			&complaint.CompanyID,
			&complaint.Title,
			&complaint.Description,
			&complaint.Active,
			&complaint.CreatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("error to scaning complaint: %v", err)
		}

		complaints = append(complaints, complaint)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over complaints: %v", err)
	}

	return complaints, nil
}

func FindComplaint(complaint models.Complaint) (*models.Complaint, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", selectComplaintFields, complaintTable)
	err := database.DB.QueryRow(context.Background(), sql, complaint.ID).Scan(
		&complaint.ID,
		&complaint.UserID,
		&complaint.CompanyID,
		&complaint.Title,
		&complaint.Description,
		&complaint.Active,
		&complaint.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error to find company: %v", err)
	}

	return &complaint, nil
}

func FindCompanyByCompanyId(complaint models.Complaint) (*models.Complaint, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE company_id = $1", selectComplaintFields, complaintTable)
	err := database.DB.QueryRow(context.Background(), sql, complaint.CompanyID).Scan(
		&complaint.ID,
		&complaint.UserID,
		&complaint.CompanyID,
		&complaint.Title,
		&complaint.Description,
		&complaint.Active,
		&complaint.CreatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("error to find company: %v", err)
	}

	return &complaint, nil
}

func CreateComplaint(complaint models.Complaint) error {
	sql := fmt.Sprintf(`
		INSERT INTO %s (%s)
		VALUES ($1, $2, $3, $4, $5, $6)
		RETURNING id
	`, complaintTable, insertComplaintFields)

	err := database.DB.QueryRow(context.Background(), sql,
		complaint.UserID,
		complaint.CompanyID,
		complaint.Title,
		complaint.Description,
		&complaint.Active,
		time.Now(),
	).Scan(&complaint.ID)

	if err != nil {
		return fmt.Errorf("error to insert complaint: %v", err)
	}

	return nil
}

func UpdateComplaint(complaint models.Complaint) error {
	var setClauses []string
	var args []interface{}
	argIndex := 1

	if complaint.Title != "" {
		setClauses = append(setClauses, fmt.Sprintf("title = $%d", argIndex))
		args = append(args, complaint.Title)
		argIndex++
	}

	if complaint.Description != "" {
		setClauses = append(setClauses, fmt.Sprintf("description = $%d", argIndex))
		args = append(args, complaint.Description)
		argIndex++
	}

	if len(setClauses) == 0 {
		return fmt.Errorf("no fields to update")
	}

	sql := fmt.Sprintf(`
		UPDATE %s
		SET %s
		WHERE id = $%d
		RETURNING id
	`, complaintTable, strings.Join(setClauses, ", "), argIndex)

	args = append(args, complaint.ID)

	err := database.DB.QueryRow(context.Background(), sql, args...).Scan(&complaint.ID)
	if err != nil {
		return fmt.Errorf("error updating complaint: %v", err)
	}

	return nil
}

func DeleteComplaint(complaint models.Complaint) error {
	sql := fmt.Sprintf(`
		UPDATE %s
		SET active = false
		WHERE id = $1
		RETURNING id
	`, complaintTable)

	err := database.DB.QueryRow(context.Background(), sql, complaint.ID).Scan(&complaint.ID)
	if err != nil {
		return fmt.Errorf("error to delete complaint: %v", err)
	}

	return nil
}
