package services

import (
	
	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/repositories"
)

func GetAllCompanies() ([]models.Company, error) {
	companies, err := repositories.FindAllCompanies()
	if err != nil {
		return nil, err
	}

	return companies, nil
}

func GetCompany(company models.Company) (*models.Company, error) {
	foundCompany, err := repositories.FindCompany(company)
	if err != nil {
		return nil, err
	}

	return foundCompany, nil
}

func CreateCompany(company models.Company) error {
	return repositories.CreateCompany(company)
}

func UpdateCompany(company models.Company) error {
	_, err := repositories.FindCompany(company)
	if err != nil {
		return err
	}
	
	return repositories.UpdateCompany(company)
}

func DeleteCompany(id int) error {
	var company models.Company = models.Company{ID: id}

	foundCompany, err := repositories.FindCompany(company)
	if err != nil {
		return err
	}

	return repositories.DeleteCompany(*foundCompany)
}