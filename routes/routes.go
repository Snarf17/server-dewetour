package routes

import "github.com/gorilla/mux"

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	TripRoutes(r)
	CountryRoutes(r)
	TransactionRoutes(r)
}
