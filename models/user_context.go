package models

type UserContext struct {
	Username string   `json:"username"`
	Roles    []string `json:"roles"`
}
