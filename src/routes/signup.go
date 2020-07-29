package routes

import (
	"fmt"
	"net/http"
)

func SignUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	fmt.Fprint(w, "Sign Up")
}
