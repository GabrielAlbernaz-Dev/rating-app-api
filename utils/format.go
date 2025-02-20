package utils

import (
	"regexp"
)

func FormatCpfCnpj(cnpj string) string {
	return regexp.MustCompile(`\D`).ReplaceAllString(cnpj, "")
}

func FormatPhone(phone string) string {
	return regexp.MustCompile(`\D`).ReplaceAllString(phone, "")
}