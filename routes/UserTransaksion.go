package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gorilla/mux"
)

func UserTransactionRoutes(r *mux.Router) {
	TransRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransRepository)

	r.HandleFunc("/transaction", h.FindTransaction).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(middleware.UploadFile(h.CreateTransaction))).Methods("POST")
	// r.HandleFunc("/transaction/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTransaction))).Methods("PATCH")
	// r.HandleFunc("/users/{id}", h.UpdateUsers).Methods("PATCH")
	// r.HandleFunc("/users/{id}", h.DeleteUsers).Methods("DELETE")

}
