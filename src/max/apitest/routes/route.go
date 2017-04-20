package routes

import (
	"max/apitest/rest"
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
	Route{
		"root",
		"GET",
		"/",
		rest.Dummy,
	},
	Route{
		"list",
		"GET",
		"/private/list",
		rest.Validate(rest.List),
	},
	Route{
		"select",
		"GET",
		"/private/select/{userId}",
		rest.Validate(rest.Select),
	},
	Route{
		"simulaAutenticazione",
		"GET",
		"/settoken",
		rest.SetToken,
	},
	Route{
		"private",
		"GET",
		"/private",
		rest.Validate(rest.ProtectedProfile),
	},
	Route{
		"logout",
		"GET",
		"/logout",
		rest.Logout,
	},
}
