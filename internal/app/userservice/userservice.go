package userservice

import (
	"Strooer/internal/app/model"
	"log"
	"net/http"
)

type Application struct {
}

type UserApi struct {}

type UserService struct {
	UserConnection interface{
		FetchUsers(*int, chan []model.User)
		FetchComments(*int, chan []model.Comment)
		DoRequest(req *http.Request) ([]byte, error)
	}
}

func (app *Application) Start() {
	router := app.NewRouter()
	port := "8081"

	log.Printf("NOTICE: starting server on port %s", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
