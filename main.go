package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Tasks struct {
	ID         string `json:"id"`
	TaskName   string `json:"taskName"`
	TaskDetail string `json:"taskDetail"`
	Date       string `json:"date"`
}

var tasks []Tasks

func allTasks() {
	task := Tasks{
		ID:         "1",
		TaskName:   "New projects",
		TaskDetail: "You must lead the project and finish it",
		Date:       "2022-01-22",
	}
	tasks = append(tasks, task)
	task1 := Tasks{
		ID:         "2",
		TaskName:   "Power project",
		TaskDetail: "We need to hire more staffs before the deadline",
		Date:       "2022-01-22",
	}
	tasks = append(tasks, task1)
	fmt.Println("Your tasks are", tasks)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am at the home page")
}

func getAllTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}
func task(w http.ResponseWriter, r *http.Request) {
	taskId := mux.Vars(r)
	for i := 0; i < len(tasks); i++ {
		if taskId["id"] == tasks[i].ID {
			json.NewEncoder(w).Encode(tasks[i])
			return

		}
	}
	// if we do not find the task
	json.NewEncoder(w).Encode(map[string]string{"status": "Error"})
}
func createTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am at the home page")
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am at the home page")
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am at the home page")
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("Get")
	router.HandleFunc("/gettasks", getAllTasks).Methods("Get")
	router.HandleFunc("/gettask/{id}", task).Methods("Get")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", router))
}

func main() {
	allTasks()
	fmt.Println("Hello Fulutter boys")
	handleRoutes()

}
