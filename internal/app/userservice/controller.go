package userservice

import (
	"Strooer/internal/app/model"
	"encoding/json"
	"log"
	"net/http"
)

func (app *Application) HealthCheck(w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	resp := model.HealthResponse{
		Message: "healthy",
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Printf("ERROR: could not marshal health response to json with err: %v", err)
		app.HttpResponse(w, http.StatusInternalServerError, []byte(err.Error()))
		return
	}

	app.HttpResponse(w, http.StatusOK, jsonResp)
}



func (app *Application) HttpResponse(w http.ResponseWriter, code int, response []byte) {
	w.WriteHeader(code)
	_, _ = w.Write(response)
}