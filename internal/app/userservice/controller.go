package userservice

import (
	"Strooer/internal/app/model"
	"encoding/json"
	"log"
	"net/http"
)

func (app *Application) GetUsersWithComments(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	us := app.NewUserService()
	result, err := us.GetUsersWithComments(nil)

	if err != nil {
		resp, err := json.Marshal(model.ApiError{Code: http.StatusInternalServerError, Message: err.Error()})
		if err != nil {
			log.Printf("ERROR: could not marshal api error with error: %v", err)
			app.HttpResponse(w, http.StatusInternalServerError, []byte("unexpected error"))
			return
		}
		app.HttpResponse(w, http.StatusInternalServerError, resp)
		return
	}
	successResp := model.UsersAndCommentsResponse{Result: result}
	resp, err := json.Marshal(successResp)
	if err != nil {
		log.Printf("ERROR: could not marshal get users success message with error: %v", err)
		app.HttpResponse(w, http.StatusInternalServerError, []byte("unexpected error"))
		return
	}

	app.HttpResponse(w, http.StatusOK, resp)
}

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
