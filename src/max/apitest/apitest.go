package main

import (
	"log"
	"max/apitest/routes"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8080", routes.NewRouter()))
}
