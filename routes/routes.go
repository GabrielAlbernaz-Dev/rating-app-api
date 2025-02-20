package routes

import (
	"log"
	"net/http"

	"github.com/gabrielalbernazdev/rating-app-api/controllers"
	"github.com/gabrielalbernazdev/rating-app-api/middlewares"
	"github.com/gorilla/mux"
)

func Init() {
	router := mux.NewRouter()
	router.Use(middlewares.JsonMiddleware)

	apiRouter := router.PathPrefix("/api").Subrouter()

	protectedApiRouter := router.PathPrefix("/api").Subrouter()
	protectedApiRouter.Use(middlewares.AuthMiddleware)

	apiRouter.HandleFunc("/auth/login", controllers.AuthLogin).Methods("POST")
	apiRouter.HandleFunc("/auth/register", controllers.AuthRegister).Methods("POST")

	protectedApiRouter.HandleFunc("/companies", controllers.GetAllCompanies).Methods("GET")
	protectedApiRouter.HandleFunc("/companies/{id}", controllers.GetCompany).Methods("GET")
	protectedApiRouter.HandleFunc(
		"/companies", middlewares.HasAnyRole([]string{"ADMIN"})(controllers.CreateCompany),
	).Methods("POST")
	protectedApiRouter.HandleFunc(
		"/companies/{id}", middlewares.HasAnyRole([]string{"ADMIN"})(controllers.UpdateCompany),
	).Methods("PUT")
	protectedApiRouter.HandleFunc(
		"/companies/{id}", middlewares.HasAnyRole([]string{"ADMIN"})(controllers.DeleteCompany),
	).Methods("DELETE")
	protectedApiRouter.HandleFunc("/complaints", controllers.GetAllComplaints).Methods("GET")
	protectedApiRouter.HandleFunc("/complaints/{id}", controllers.GetComplaint).Methods("GET")
	protectedApiRouter.HandleFunc("/complaints", controllers.CreateComplaint).Methods("POST")
	protectedApiRouter.HandleFunc("/complaints/{id}", controllers.UpdateComplaint).Methods("PUT")
	protectedApiRouter.HandleFunc("/complaints/{id}", controllers.DeleteComplaint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))
}