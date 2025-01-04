package routes

import (
	"github.com/themrgeek/cleaning-service/pkg/controllers"

	"github.com/gorilla/mux"
)

func SetupRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/signup", controllers.Signup).Methods("POST")
	r.HandleFunc("/login", controllers.Login).Methods("POST")

	userRoutes := r.PathPrefix("/user").Subrouter()
	userRoutes.HandleFunc("/book-service", controllers.BookService).Methods("POST")
	userRoutes.HandleFunc("/verify-otp", controllers.VerifyOTP).Methods("POST")
	userRoutes.HandleFunc("/payment", controllers.ProcessPayment).Methods("POST")
	userRoutes.HandleFunc("/review", controllers.SubmitReview).Methods("POST")
	userRoutes.HandleFunc("/delete-appointment/{id}", controllers.DeleteAppointment).Methods("DELETE")

	cleanerRoutes := r.PathPrefix("/cleaner").Subrouter()
	cleanerRoutes.HandleFunc("/status", controllers.UpdateStatus).Methods("POST")
	cleanerRoutes.HandleFunc("/inquiries", controllers.ViewInquiries).Methods("GET")
	cleanerRoutes.HandleFunc("/inquiries/{id}/complete", controllers.CompleteInquiry).Methods("POST")

	adminRoutes := r.PathPrefix("/admin").Subrouter()
	adminRoutes.HandleFunc("/performance", controllers.ViewPerformance).Methods("GET")
	adminRoutes.HandleFunc("/revenue", controllers.ViewRevenue).Methods("GET")

	return r
}
