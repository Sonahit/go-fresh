package middleware

import "net/http"

type Middleware func(http.HandlerFunc) http.HandlerFunc

func Middlewares(handler func(w http.ResponseWriter, r *http.Request), middlewares ...Middleware) func(w http.ResponseWriter, r *http.Request) {
	if len(middlewares) < 1 {
		return handler
	}

	wrapped := handler

	for _, m := range middlewares {
		wrapped = m(wrapped)
	}

	return wrapped
}
