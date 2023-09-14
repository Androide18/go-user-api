package main

import (
	"log"
	"net/http"

	"github.com/androide18/go-user-api/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterUsersRoutes(router)
	http.Handle("/", router)
	log.Fatal(http.ListenAndServe("localhost:9010", router))
}
