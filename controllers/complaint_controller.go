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

func GetAllComplaints(w http.ResponseWriter, r *http.Request) {
	complaints, err := services.GetAllComplaints()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()})
		return
	}

	var response dtos.ListResponseDto

	if len(complaints) > 0 {
		response = dtos.ListResponseDto{List: []interface{}{complaints}}
	} else {
		response = dtos.ListResponseDto{List: make([]interface{}, 0)}
	}

	utils.WriteJson(
		w,
		response,
		http.StatusOK,
	)
}

func GetComplaint(w http.ResponseWriter, r *http.Request) {
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

	complaint := models.Complaint{ID: id}
	complaintResponse, err := services.GetComplaint(complaint)
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
		dtos.ItemResponseDto{Item: complaintResponse},
		http.StatusOK,
	)
}

func CreateComplaint(w http.ResponseWriter, r *http.Request) {
	var complaint models.Complaint

	if err := json.NewDecoder(r.Body).Decode(&complaint); err != nil {
		utils.WriteJson(
			w,
			dtos.GenericErrorResponseDto{Error: "Invalid JSON format: " + err.Error(), Timestamp: time.Now()},
			http.StatusBadRequest,
		)
		return
	}

	if err := validations.ValidateComplaintCreateBody(complaint); err != nil {
		utils.WriteJson(
			w,
			dtos.GenericErrorResponseDto{Error: "Validation failed: " + err.Error(), Timestamp: time.Now()},
			http.StatusUnprocessableEntity,
		)
		return
	}

	if err := services.CreateComplaint(complaint); err != nil {
		utils.WriteJson(
			w,
			dtos.GenericErrorResponseDto{Error: "Failed to create complaint: " + err.Error(), Timestamp: time.Now()},
			http.StatusInternalServerError,
		)
		return
	}

	utils.WriteJson(
		w,
		dtos.MessageResponseDto{Message: "Complaint successfully created"},
		http.StatusCreated,
	)
}

func UpdateComplaint(w http.ResponseWriter, r *http.Request) {
    var complaint models.Complaint
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

    if err = json.NewDecoder(r.Body).Decode(&complaint); err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
    }

    complaint.ID = id

	err = validations.ValidateComplaintUpdateBody(complaint)
	if err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusUnprocessableEntity,
        )
        return
	}

    if err = services.UpdateComplaint(complaint); err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
    }

    utils.WriteJson(
        w,
        dtos.MessageResponseDto{Message: "Complaint successfully updated"},
        http.StatusOK,
    )
}


func DeleteComplaint(w http.ResponseWriter, r *http.Request) {
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

	if err := services.DeleteComplaint(id); err != nil {
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

