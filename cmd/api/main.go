package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

const port = 8080

type Application struct {
	DSN    string
	Domain string
}

func main() {
	// set application config
	var app Application

	// read from command line
	flag.StringVar(&app.DSN, "dsn", "host=localhost port=5432 user=postgres password=postgres dbname=movies sslmode=disable timezone=UTC connect_timeout=5", "Postgres connection string")
	flag.Parse()

	// connect to the database

	app.Domain = "example.com"

	log.Printf("Starting application on port: %d", port)

	// start a web server
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), app.routes())
	if err != nil {
		log.Fatalf("We encountered an error: %v", err)
	}
}
