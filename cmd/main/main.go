package main

import (
	"log"
	"net/http"

	"github.com/alifurkangokce/GoBookStoreStudy/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9910", r))

}
