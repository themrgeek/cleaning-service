package main

import (
	"log"
	"net/http"

	"github.com/themrgeek/cleaning-service/pkg/config"
	"github.com/themrgeek/cleaning-service/pkg/routes"
)

func main() {
	config.LoadConfig()

	r := routes.SetupRouter()
	log.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
