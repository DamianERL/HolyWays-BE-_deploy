package routes

import (
	"holyways/handlers"
	"holyways/repositories"
	"holyways/pkg/mysql"
	"holyways/pkg/middleware"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router){
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h:= handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions",middleware.Auth(h.FindTransactions)).Methods("GET")
	r.HandleFunc("/transactionss/{id}",middleware.Auth(h.FindTransactionss)).Methods("GET")
	r.HandleFunc("/transactionx/{id}",middleware.Auth(h.FindTransactionx)).Methods("GET")
	r.HandleFunc("/transaction",middleware.Auth(h.CreateTransaction)).Methods("POST")
	//midtrans
	r.HandleFunc("/notification", h.Notification).Methods("POST")
}