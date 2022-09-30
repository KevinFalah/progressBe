package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/gorilla/mux"
)

func HiredRoutes(r *mux.Router) {
	HiredRepository := repositories.RepositoryHired(mysql.DB)
	h := handlers.HandlerHired(HiredRepository)

	r.HandleFunc("/hireds", h.FindHireds).Methods("GET")
	r.HandleFunc("/hired/{id}", h.GetHired).Methods("GET")
	r.HandleFunc("/hired", middleware.Auth(h.CreateHired)).Methods("POST")
	// r.HandleFunc("/hired/{id}", middleware.Auth(h.UpdateHired)).Methods("PATCH")
	// r.HandleFunc("/hired/{id}", middleware.Auth(h.DeleteHired)).Methods("DELETE")
}
