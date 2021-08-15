package main

import (
	"github.com/gorilla/mux"
	"go-todos/apis"
	"go-todos/handlers"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	client := &apis.HttpClient{}
	r.HandleFunc("/todos", handlers.GetTodos(client))

	err := http.ListenAndServe(":3333", r)
	if err != nil {
		return 
	}
}
