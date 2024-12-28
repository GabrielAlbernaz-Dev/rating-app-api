package dtos

import "time"

type TokenResponseDto struct {
	Token string `json:"token"`
}

type GenericErrorResponseDto struct {
	Error     string    `json:"error"`
	Timestamp time.Time `json:"timestamp"`
}

type ListResponseDto struct {
	List []interface{} `json:"list"`
}

type ItemResponseDto struct {
	Item interface{} `json:"item"`
}

type MessageResponseDto struct {
	Message string `json:"message"`
}
