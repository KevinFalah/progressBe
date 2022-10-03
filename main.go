package main

import (
	"fmt"
	"net/http"
	"waysgallery/database"
	"waysgallery/pkg/mysql"
	"waysgallery/routes"
	"github.com/joho/godotenv"
	"github.com/gorilla/handlers"

	"github.com/gorilla/mux"
)

func main() {

	errEnv := godotenv.Load()
	if errEnv != nil {
		panic("Failed to load env file")
	}

	// initial DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	r.PathPrefix("/uploads").Handler(http.StripPrefix("/uploads/", http.FileServer(http.Dir("./uploads"))))


	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

		// Setup allowed Header, Method, and Origin for CORS on this below code ...
		var AllowedHeaders = handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
		var AllowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "PATCH", "DELETE"})
		var AllowedOrigins = handlers.AllowedOrigins([]string{"*"})


		var port = "5000"
		fmt.Println("server running localhost:" + port)

	// Embed the setup allowed in 2 parameter on this below code
	http.ListenAndServe("localhost:"+port, handlers.CORS(AllowedHeaders, AllowedMethods, AllowedOrigins)(r))
}
