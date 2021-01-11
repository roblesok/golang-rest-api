package main

import (
	"log"
	"net/http"

	"github.com/roblesok/golang-rest-api/handler"
)

func main() {
	port := ":3000"
	ph := handler.NewBookHandler()

	http.Handle("/api/products", ph)
	http.Handle("/api/products/", ph)

	log.Fatal(http.ListenAndServe(port, nil))
}
