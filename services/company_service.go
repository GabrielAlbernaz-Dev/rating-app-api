package services

import (
	"fmt"
	
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
	existedCompany, _ := repositories.FindCompanyByCNPJ(company)
	if existedCompany != nil && existedCompany.ID != 0 {
		return fmt.Errorf("company with cnpj %s already exists", existedCompany.CNPJ)
	}

	return repositories.CreateCompany(company)
}

func UpdateCompany(company models.Company) error {
	foundCompany, _ := repositories.FindCompany(company)

	if foundCompany == nil || foundCompany.ID == 0 {
		return fmt.Errorf("company with cnpj %s not found", company.CNPJ)
	}
	
	return repositories.UpdateCompany(company)
}

func DeleteCompany(id int) error {
	var company models.Company = models.Company{ID: id}

	foundCompany, err := repositories.FindCompany(company)
	if err != nil {
		return err
	}

	if foundCompany.ID == 0 {
		return fmt.Errorf("company with cnpj %s not found", foundCompany.CNPJ)
	}

	return repositories.DeleteCompany(*foundCompany)
}