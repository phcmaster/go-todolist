package main

import (
	"net/http"
	"log"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"

)



type Todo struct {
	ID 			string 	`json:"id"`
	Item 		string 	`json:"Item"`
	Completed 	bool 	`josn:"Completed"`
}


var todos = []Todo{
	{ID: "1", Item: "Put the trash outsite", Completed: false},
	{ID: "2", Item: "Buy water", Completed: true},
	{ID: "3", Item: "study go", Completed: true},
	{ID: "4", Item: "practice exercice", Completed: false},
}


func getTodos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var newTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
		log.Fatal(err)
	}
	todos = append(todos, newTodo)

	json.NewEncoder(w).Encode(todos)
}


func update(w http.ResponseWriter, r *http.Request) {
	var updatedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&updatedTodo); err != nil {
		log.Fatal(err)
	}
	for i, t := range todos {
		if t.ID == updatedTodo.ID {
			todos[i] = updatedTodo
		}
	}
	json.NewEncoder(w).Encode(todos)
}

func delete(w http.ResponseWriter, r *http.Request) {
	var deletedTodo Todo
	if err := json.NewDecoder(r.Body).Decode(&deletedTodo); err != nil {
		log.Fatal(err)
	}
	for i, t := range todos {
		if t.ID == deletedTodo.ID {
			todos = append(todos[:i], todos[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(todos)
}


func main() {

router := mux.NewRouter()

router.HandleFunc("/todos", getTodos).Methods("GET")

router.HandleFunc("/createTodo", createTodo).Methods("POST")

router.HandleFunc("/update/{id}", update).Methods("PUT")

router.HandleFunc("/delete/{id}", delete).Methods("DELETE")

fmt.Println("Server lisening on port 8081...")
log.Fatal(http.ListenAndServe(":8081", nil))


}