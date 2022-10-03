package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	PostRoutes(r)
	HiredRoutes(r)
	FollowingRoutes(r)
}
