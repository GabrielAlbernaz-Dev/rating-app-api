package repositories

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/gabrielalbernazdev/rating-app-api/infra/database"
	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/utils"
)

const (
	companyTable = "companies"
	selectCompanyFields = "id, name, cnpj, address, email, phone, category_id, registration_date, active"
	insertCompanyFields = "name, cnpj, address, email, phone, category_id, registration_date"
)

func FindAllCompanies() ([]models.Company, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s", selectCompanyFields, companyTable)

	rows, err := database.DB.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("error to get companies: %v", err)
	}
	defer rows.Close()

	var companies []models.Company

	for rows.Next() {
		var company models.Company
		err := rows.Scan(
			&company.ID,
			&company.CNPJ,
			&company.Address,
			&company.Email,
			&company.Phone,
			&company.CategoryID,
			&company.RegistrationDate,
			&company.Active,
		)
		if err != nil {
			return nil, fmt.Errorf("error to scaning company: %v", err)
		}
		companies = append(companies, company)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating over companies: %v", err)
	}

	return companies, nil
}

func FindCompany(company models.Company) (*models.Company, error) {
	sql := fmt.Sprintf("SELECT %s FROM %s WHERE id = $1", selectCompanyFields, companyTable)
	err := database.DB.QueryRow(context.Background(), sql, company.ID).Scan(
		&company.ID, 
		&company.Name, 
		&company.CNPJ,
		&company.Address, 
		&company.Email, 
		&company.Phone,
		&company.CategoryID,
		&company.RegistrationDate,
		&company.Active,
	)
	if err != nil {
		return nil, fmt.Errorf("error to find company: %v", err)
	}

	return &company, nil
}

func CreateCompany(company models.Company) error {
	sql := fmt.Sprintf(`
		INSERT INTO companies (%s)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id
	`, insertCompanyFields)

	company.CNPJ = utils.FormatCpfCnpj(company.CNPJ)

	err := database.DB.QueryRow(context.Background(), sql, company.Name, company.CNPJ, company.Address, company.Email, company.Phone, company.CategoryID, time.Now()).Scan(&company.ID)
	if err != nil {
		return fmt.Errorf("error to insert company: %v", err)
	}

	return nil
}

func UpdateCompany(company models.Company) error {
	var setClauses []string
	var args []interface{}
	argIndex := 1

	if company.Name != "" {
		setClauses = append(setClauses, fmt.Sprintf("name = $%d", argIndex))
		args = append(args, company.Name)
		argIndex++
	}

	if company.CNPJ != "" {
		setClauses = append(setClauses, fmt.Sprintf("cnpj = $%d", argIndex))
		args = append(args, company.CNPJ)
		argIndex++
	}

	if company.Address != "" {
		setClauses = append(setClauses, fmt.Sprintf("address = $%d", argIndex))
		args = append(args, company.Address)
		argIndex++
	}

	if company.Email != "" {
		setClauses = append(setClauses, fmt.Sprintf("email = $%d", argIndex))
		args = append(args, company.Email)
		argIndex++
	}

	if company.Phone != "" {
		setClauses = append(setClauses, fmt.Sprintf("phone = $%d", argIndex))
		args = append(args, company.Phone)
		argIndex++
	}

	if company.CategoryID != nil {
		setClauses = append(setClauses, fmt.Sprintf("category_id = $%d", argIndex))
		args = append(args, *company.CategoryID)
		argIndex++
	}

	if !company.RegistrationDate.IsZero() {
		setClauses = append(setClauses, fmt.Sprintf("registration_date = $%d", argIndex))
		args = append(args, company.RegistrationDate)
		argIndex++
	}

	sql := fmt.Sprintf(`
		UPDATE %s
		SET %s
		WHERE id = $%d
		RETURNING id
	`, companyTable, strings.Join(setClauses, ", "), argIndex)

	args = append(args, company.ID)

	err := database.DB.QueryRow(context.Background(), sql, args...).Scan(&company.ID)
	if err != nil {
		return fmt.Errorf("error to update company: %v", err)
	}

	return nil
}

func SetCompanyActive(company models.Company) error {
	sql := fmt.Sprintf(`
		UPDATE %s
		SET active = true
		WHERE id = $1
		RETURNING id
	`, companyTable)

	err := database.DB.QueryRow(context.Background(), sql, company.ID).Scan(&company.ID)
	if err != nil {
		return fmt.Errorf("error to activate company: %v", err)
	}

	return nil
}

func DeleteCompany(company models.Company) error {
	sql := fmt.Sprintf(`
		UPDATE %s
		SET active = false
		WHERE id = $1
		RETURNING id
	`, companyTable)

	err := database.DB.QueryRow(context.Background(), sql, company.ID).Scan(&company.ID)
	if err != nil {
		return fmt.Errorf("error to delete company: %v", err)
	}

	return nil
}
