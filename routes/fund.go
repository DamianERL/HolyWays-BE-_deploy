package routes

import (
	"holyways/handlers"
	"holyways/pkg/middleware"
	"holyways/pkg/mysql"
	"holyways/repositories"

	"github.com/gorilla/mux"
)

func FundRoutes(r *mux.Router)  {
	fundRepository := repositories.RepositoryFund(mysql.DB)
	h := handlers.HandlerFund(fundRepository)

	r.HandleFunc("/funds",h.FindFunds).Methods("GET")
	r.HandleFunc("/fund/{id}",middleware.Auth(h.GetFund)).Methods("GET")
	r.HandleFunc("/fundss",middleware.Auth(h.FindFundId)).Methods("GET")
	r.HandleFunc("/fund",middleware.Auth(middleware.UploadFile(h.CreateFund))).Methods("POST")
	r.HandleFunc("/fund/{id}",middleware.Auth(middleware.UploadFile(h.UpdateFund))).Methods("PACTH")
	r.HandleFunc("/fund/{id}",middleware.Auth(h.DeleteFund)).Methods("DELETE")
}