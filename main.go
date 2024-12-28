package main

import (
	"github.com/gabrielalbernazdev/rating-app-api/infra/config"
	"github.com/gabrielalbernazdev/rating-app-api/infra/database"
	"github.com/gabrielalbernazdev/rating-app-api/routes"
)

type ContextKey string
const CurrentUserKey ContextKey = "currentUser"

func main() {
	config.LoadEnv()

	database.Connect()

	routes.Init()
}