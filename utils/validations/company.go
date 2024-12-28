package validations

import (

	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/utils"
)

func ValidateCompanyCreateBody(company models.Company) error {
	if company.Name == "" {
		return utils.ErrorField("name","required")
	}

	if company.CNPJ == "" {
		return utils.ErrorField("CNPJ","required")
	}

	if !utils.ValidateCNPJ(company.CNPJ) {
		return utils.ErrorField("CNPJ","invalid")
	}

	return nil
}

func ValidateCompanyUpdateBody(company models.Company) error {
	if company.CNPJ != "" && !utils.ValidateCNPJ(company.CNPJ) {
		return utils.ErrorField("CNPJ","invalid")
	}

	return nil
}