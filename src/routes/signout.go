package routes

import (
	"fmt"
	"net/http"
)

func SignOut(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	fmt.Fprint(w, "SignOut")
}
