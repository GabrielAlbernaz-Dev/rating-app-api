package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gabrielalbernazdev/rating-app-api/dtos"
	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/services"
	"github.com/gabrielalbernazdev/rating-app-api/utils"
	"github.com/gabrielalbernazdev/rating-app-api/utils/validations"
	"github.com/gorilla/mux"
)

func GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	companies, err := services.GetAllCompanies()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()})
		return
	}

	var response dtos.ListResponseDto

	if len(companies) > 0 {
		response = dtos.ListResponseDto{List: []interface{}{companies}}
	} else {
		response = dtos.ListResponseDto{List: make([]interface{}, 0)}
	}
	
	utils.WriteJson(
		w,
		response,
		http.StatusOK,
	)
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	vars :=  mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
	}
	
	company := models.Company{ID: id}
	companyResponse, err := services.GetCompany(company)
	if err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusUnprocessableEntity,
        )
        return
	}

	utils.WriteJson(
		w,
		dtos.ItemResponseDto{Item: companyResponse},
		http.StatusOK,
	)
}

func CreateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	var err error

	if err = json.NewDecoder(r.Body).Decode(&company); err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
	}

    err = validations.ValidateCompanyCreateBody(company)
	if err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusUnprocessableEntity,
        )
        return
	}

	if err = services.CreateCompany(company); err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
	}

	utils.WriteJson(
		w,
		dtos.MessageResponseDto{Message: "Company successfully created"},
		http.StatusCreated,
	)
}

func UpdateCompany(w http.ResponseWriter, r *http.Request) {
	var company models.Company
	var err error

	vars := mux.Vars(r)
	id, err := strconv.Atoi(strings.TrimSpace(vars["id"]))
	if err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
	}

	if err = json.NewDecoder(r.Body).Decode(&company); err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
	}

	company.ID = id

	err = validations.ValidateCompanyCreateBody(company)
	if err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusUnprocessableEntity,
        )
        return
	}

	if err = services.UpdateCompany(company); err != nil {
		utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
	}

	utils.WriteJson(
		w,
		dtos.MessageResponseDto{Message: "Company successfully updated"},
		http.StatusOK,
	)
}

func DeleteCompany(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
	}

	if err := services.DeleteCompany(id); err != nil {
		utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
	}

	utils.WriteJson(
		w,
		nil,
		http.StatusNoContent,
	)
}
