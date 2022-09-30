package routes

import (
	"waysgallery/handlers"
	"waysgallery/pkg/middleware"
	"waysgallery/pkg/mysql"
	"waysgallery/repositories"

	"github.com/gorilla/mux"
)

func PostRoutes(r *mux.Router) {
	PostRepository := repositories.RepositoryPost(mysql.DB)
	h := handlers.HandlerPost(PostRepository)

	r.HandleFunc("/posts", h.FindPosts).Methods("GET")
	r.HandleFunc("/post/{id}", h.GetPost).Methods("GET")
	r.HandleFunc("/post", middleware.Auth(middleware.UploadFile(h.CreatePost))).Methods("POST")
	// r.HandleFunc("/post/{id}", middleware.Auth(h.UpdatePost)).Methods("PATCH")
	// r.HandleFunc("/post/{id}", middleware.Auth(h.DeletePost)).Methods("DELETE")
}
