package routes

import (
	"dewetour/handlers"
	"dewetour/pkg/middleware"
	"dewetour/pkg/mysql"
	"dewetour/repositories"

	"github.com/gorilla/mux"
)

func TripRoutes(r *mux.Router) {
	TripRepository := repositories.RepositoryTrip(mysql.DB)
	h := handlers.HandlerTrips(TripRepository)

	r.HandleFunc("/trips", h.FindTrip).Methods("GET")
	r.HandleFunc("/trips/{id}", h.GetTrip).Methods("GET")
	r.HandleFunc("/trips", middleware.AuthAdmin(middleware.UploadFile(h.CreateTrip))).Methods("POST")
	r.HandleFunc("/trips/{id}", middleware.AuthAdmin(middleware.UploadFile(h.UpdateTrip))).Methods("PATCH")
	r.HandleFunc("/trips/{id}", h.DeleteTrip).Methods("DELETE")

}
