package handlers

import (
	"encoding/json"
	"go-todos/apis"
	"go-todos/models"
	"net/http"
)

var URL = "https://jsonplaceholder.typicode.com/todos"

func GetTodos(client apis.HttpInterface) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request) {
		var todos []models.Todo

		body, err := client.Get(URL)
		returnError(w, err)

		err = json.Unmarshal(body, &todos)
		returnError(w, err)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(todos)
		if err != nil {
			return
		}

	}
}



func returnError(w http.ResponseWriter, err error) {
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
}