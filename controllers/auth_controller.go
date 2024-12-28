package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gabrielalbernazdev/rating-app-api/dtos"
	"github.com/gabrielalbernazdev/rating-app-api/models"
	"github.com/gabrielalbernazdev/rating-app-api/services"
	"github.com/gabrielalbernazdev/rating-app-api/utils"
	"github.com/gabrielalbernazdev/rating-app-api/utils/validations"
)

func AuthLogin(w http.ResponseWriter, r *http.Request) {
    var user models.User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
    }

	if validationErr := validations.ValidateUserLoginBody(user); validationErr != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: validationErr.Error(), Timestamp: time.Now()},
            http.StatusUnprocessableEntity,
        )
        return
	}

    login, loginErr := services.AuthVerifyUser(user)
    if loginErr != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: loginErr.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
    }

    if !login {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: "Invalid Credentials", Timestamp: time.Now()},
            http.StatusUnauthorized,
        )
        return
    }

    token, err := services.GenerateToken(user.Email)
    if err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: "Error to generate token", Timestamp: time.Now()},
            http.StatusInternalServerError,
        )
        return
    }

    utils.WriteJson(
        w,
        dtos.TokenResponseDto{Token: token},
        http.StatusOK,
    )
}

func AuthRegister(w http.ResponseWriter, r *http.Request) {
    var user models.User

    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: err.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
    }

	if validationErr := validations.ValidateUserRegisterBody(user); validationErr != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: validationErr.Error(), Timestamp: time.Now()},
            http.StatusUnprocessableEntity,
        )
        return
	}

    if registerErr := services.AuthRegisterUser(user); registerErr != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: registerErr.Error(), Timestamp: time.Now()},
            http.StatusBadRequest,
        )
        return
    }

    token, err := services.GenerateToken(user.Email)
    if err != nil {
        utils.WriteJson(
            w,
            dtos.GenericErrorResponseDto{Error: "Error to generate token", Timestamp: time.Now()},
            http.StatusInternalServerError,
        )
        return
    }

    utils.WriteJson(
        w,
        dtos.TokenResponseDto{Token: token},
        http.StatusOK,
    )
}