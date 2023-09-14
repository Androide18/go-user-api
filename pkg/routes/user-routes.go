package routes

import (
	"github.com/androide18/go-user-api/pkg/controllers"
	"github.com/gorilla/mux"
)

var RegisterUsersRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/users/", controllers.CreateUser).Methods("POST")
	router.HandleFunc("/api/users/", controllers.GetUsers).Methods("GET")
	router.HandleFunc("/api/users/{userId}", controllers.GetUserById).Methods("GET")
	router.HandleFunc("/api/users/{userId}", controllers.UpdateUser).Methods("PUT")
	router.HandleFunc("/api/users/{userId}", controllers.DeleteUser).Methods("DELETE")
}
