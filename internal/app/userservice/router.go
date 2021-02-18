package userservice

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}
type Routes []Route

const (
	ApiPattern = "/api/v1"
	UserApiPattern = "/user"
	UserWithCommentsApiPattern = "/users-with-comments"
	HealthCheckPattern = "/health"
)

func (app *Application) NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(false)
	for _, route := range app.routes() {
		handlerFunc := route.HandlerFunc
		var r = new(mux.Route)
		r = router.Handle(route.Pattern, app.LoggerMiddleware(route.Name, handlerFunc))
		r.Name(route.Name).Methods(route.Method)
	}
	return router
}

// custom logger middleware
func (app *Application) LoggerMiddleware(routeName string, next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s %s", r.Method, r.RequestURI, routeName, time.Since(time.Now()))
		next(w, r)
	})
}

func (app *Application) routes() Routes {
	return Routes{
		Route{
			"HealthCheck",
			http.MethodGet,
			fmt.Sprintf("%s%s", ApiPattern, HealthCheckPattern),
			app.HealthCheck,
		},
		Route{
			"GetUsersWithComments",
			http.MethodGet,
			fmt.Sprintf("%s%s", ApiPattern, UserWithCommentsApiPattern),
			app.GetUsersWithComments,
		},
	}
}
