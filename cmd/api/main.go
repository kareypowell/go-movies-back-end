package main

import (
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type Application struct {
	Domain string
}

func main() {
	// set application config
	var app Application

	// read from command line

	// connect to the database

	app.Domain = "example.com"

	log.Printf("Starting application on port: %d", port)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatalf("We encountered an error: %v", err)
	}
}
