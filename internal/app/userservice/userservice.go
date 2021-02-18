package userservice

import (
	"log"
	"net/http"
)

type Application struct {
}

func (app *Application) Start() {
	router := app.NewRouter()
	port := "8081"

	log.Printf("NOTICE: starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}