package main

import (
	"log"
	"net/http"

	"cleaning-service/pkg/config"
	"cleaning-service/pkg/routes"
)

func main() {
	config.LoadConfig()

	r := routes.SetupRouter()
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
