package routes

import (
	"fmt"
	"net/http"
)

func SignIn(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	fmt.Fprint(w, "SignIn")
}
