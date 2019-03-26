package router

import (
	"net/http"
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	"../handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	{
		"route",
		"GET",
		"/api/building",
		handlers.BuildingGet,
	},
	{
		"route",
		"GET",
		"/api/person",
		handlers.PersonGet,
	},
	{
		"route",
		"POST",
		"/api/person",
		handlers.PersonAdd,
	},
	{
		"route",
		"DELETE",
		"/api/person/{id}",
		handlers.PersonRemove,
	},
	{
		"route",
		"PUT",
		"/api/building",
		handlers.BuildingUpdate,
	},
	{
		"route",
		"GET",
		"/api/summary",
		handlers.SummaryGet,
	},
	{
		"route",
		"GET",
		"/api/crime",
		handlers.CrimeGet,
	},
	{
		"route",
		"POST",
		"/api/crime",
		handlers.CrimeAdd,
	},

}

func NewRouter() http.Handler {
	router := mux.NewRouter().StrictSlash(true)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"}, // All origins
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
	})
	router.PathPrefix("/public/js/").Handler(http.StripPrefix("/public/js/", http.FileServer(http.Dir("./public/js/"))))
	router.PathPrefix("/public/css/").Handler(http.StripPrefix("/public/css/", http.FileServer(http.Dir("./public/css/"))))
	router.PathPrefix("/public/uploads/").Handler(http.StripPrefix("/public/uploads/", http.FileServer(http.Dir("./public/uploads/"))))
	router.PathPrefix("/public/images/").Handler(http.StripPrefix("/public/images/", http.FileServer(http.Dir("./public/images/"))))
	router.NotFoundHandler = http.HandlerFunc(handlers.NotFound)
	for _, route := range routes {
		var handler http.Handler
		handler = route.HandlerFunc
		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)
	}
	return c.Handler(router)
}
