package main

import (
	database "dewetour/databases"
	"dewetour/pkg/mysql"
	"dewetour/routes"
	"fmt"
	"net/http"
	"os"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Initial DB
	mysql.DatabaseInit()

	// Run Migration
	database.RunMigration()

	// r := mux.NewRouter()
	r := mux.NewRouter()

	// routes.Roue
	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	// Allowed header, Method , and Origins
	var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	var AllowedMethods = handlers.AllowedMethods([]string{"POST", "GET", "PATCH", "PUT", "DELETE", "OPTION", "HEAD"})
	var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})
	
	var PORT = os.Getenv('PORT')
	//path file
	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))

	fmt.Println("Server Running in localhost:9000")
	http.ListenAndServe(":"+PORT, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))

}
