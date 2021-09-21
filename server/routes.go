package server

import (
	"github.com/annakallo/travel-log-server/server/api"
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
	//Route{"Expenses", http.MethodGet, "/api/expenses", api.Expenses},
}
