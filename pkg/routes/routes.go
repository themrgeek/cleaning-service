package routes

import (
	"github.com/gorilla/mux"
	"github.com/themrgeek/cleaning-service/pkg/controllers"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	userRoutes := r.PathPrefix("/user").Subrouter()

	userRoutes.HandleFunc("/profile", controllers.Profile).Methods("GET")
	return r
}
