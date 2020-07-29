package middleware

import "net/http"

func Authorization(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return handler
}
