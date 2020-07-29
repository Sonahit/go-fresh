package routes

import (
	http "net/http"
)

type Route struct {
	pattern string
	handler http.HandlerFunc
}

var routes = []Route{
	{
		pattern: "/todo",
		handler: TodoHandler,
	},
	{
		pattern: "/signin",
		handler: SignIn,
	},
	{
		pattern: "/signout",
		handler: SignOut,
	},
	{
		pattern: "/signup",
		handler: SignUp,
	},
}

func Routes() {
	for _, h := range routes {
		http.HandleFunc(h.pattern, h.handler)
	}
}
