package server

import (
	"github.com/annakallo/travelog/server/api"
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{"Index", http.MethodGet, "/", api.Index},
	Route{"Countries", http.MethodGet, "/api/countries", api.Countries},
	Route{"CountryNew", http.MethodPost, "/api/countries", api.CountryNew},
	Route{"CountryUpdate", http.MethodPut, "/api/countries/{id}", api.CountryUpdate},
	Route{"CountryDelete", http.MethodDelete, "/api/countries/{id}", api.CountryDelete},
}
