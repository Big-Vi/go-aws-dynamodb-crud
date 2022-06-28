package main

import (
	"log"
	"net/http"

	"github.com/big-vi/go-aws-dynamodb-crud/internal/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8000", r))
}
