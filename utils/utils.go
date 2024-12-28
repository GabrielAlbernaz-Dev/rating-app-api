package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func WriteJson(w http.ResponseWriter, body any, statusCode int) {
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(body)
}

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func ErrorField(field string, constraint string) error {
	return fmt.Errorf("%v is %v", field, constraint)
}

func ValidateCNPJ(cnpj string) bool {
	cnpj = FormatCpfCnpj(cnpj)

	if len(cnpj) != 14 {
		return false
	}

	repeated := true
	for i := 1; i < len(cnpj); i++ {
		if cnpj[i] != cnpj[0] {
			repeated = false
			break
		}
	}
	if repeated {
		return false
	}

	d1 := calculateVerificationDigit(cnpj[:12], []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})
	d2 := calculateVerificationDigit(cnpj[:13], []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2})

	return cnpj[12] == d1 && cnpj[13] == d2
}

func calculateVerificationDigit(cnpjBase string, multipliers []int) byte {
	var sum int

	for i := 0; i < len(cnpjBase); i++ {
		num, _ := strconv.Atoi(string(cnpjBase[i]))
		sum += num * multipliers[i]
	}

	mod := sum % 11
	if mod < 2 {
		return '0'
	}
	return byte(48 + (11 - mod))
}