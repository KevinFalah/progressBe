package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/gorilla/mux"
)

func FollowingRoutes(r *mux.Router) {
	followingRepository := repositories.RepositoryFollowing(mysql.DB)
	h := handlers.HandlerFollowing(followingRepository)

	r.HandleFunc("/followings", h.FindFollowings).Methods("GET")
	r.HandleFunc("/following/{id}", h.GetFollowing).Methods("GET")
	r.HandleFunc("/following", h.CreateFollowing).Methods("POST")
	// r.HandleFunc("/following/{id}", h.DeleteFollowing).Methods("DELETE")
}
