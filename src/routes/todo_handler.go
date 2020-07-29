package routes

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	db "../database"
	rHttp "../http"
)

type Todo = db.Todo

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		addTodo(w, r)
	case "GET":
		getTodo(w, r)
	default:
		http.Error(w, "Not Found", 404)
	}

}

func getTodo(w http.ResponseWriter, r *http.Request) {
	log.Print("Handling GET /todo")
	response := rHttp.Response{Status: true}
	response.WriteJson(db.Todos, w)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	response := rHttp.Response{Status: true}

	todo := Todo{}

	if err := decoder.Decode(&todo); err != nil {
		response.Status = false
		response.WriteJson(fmt.Sprintf("Wrong json data: %+v", err), w)
		return
	}

	db.Todos = append(db.Todos, todo)
	log.Print(db.Todos)
	log.Print("Handling POST /todo")

	response.WriteJson(db.Todos, w)
}
