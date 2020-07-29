package main

import (
	"net/http"

	r "./routes"
)

func main() {
	r.Routes()
	http.ListenAndServe(":8080", nil)
}
