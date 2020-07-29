package database

type Status int

const (
	DONE Status = iota
	OPEN
)

type Todo struct {
	Name   string `json:"name"`
	Status Status `json:"status"`
}

var Todos = []Todo{}
