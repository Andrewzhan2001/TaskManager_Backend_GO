package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type Task struct {
	ID         string `json:"id"`
	TaskName   string `json:"taskName"`
	TaskDetail string `json:"taskDetail"`
	Date       string `json:"date"`
}

var tasks []Task

func allTasks() {
	task := Task{
		ID:         "1",
		TaskName:   "New projects",
		TaskDetail: "You must lead the project and finish it",
		Date:       "2022-01-22",
	}
	tasks = append(tasks, task)
	task1 := Task{
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
	w.Header().Set("Content-Type", "application/json")
	var task Task
	_ = json.NewDecoder(r.Body).Decode(&task)
	task.ID = strconv.Itoa(rand.Intn(1000))
	task.Date = time.Now().Format("2006-01-02")
	tasks = append(tasks, task)
	json.NewEncoder(w).Encode(tasks)
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	check := false
	for i, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:i], tasks[i+1:]...)
			check = true
			return
		}
	}
	if !check {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error in deleting task"})
	}
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	check := false
	for i, item := range tasks {
		if item.ID == params["id"] {
			tasks = append(tasks[:i], tasks[i+1:]...)
			var task Task
			_ = json.NewDecoder(r.Body).Decode(&task)
			task.ID = params["id"]
			task.Date = time.Now().Format("2006-01-02")
			tasks = append(tasks, task)
			check = true
			json.NewEncoder(w).Encode(tasks)
			return
		}
	}
	if !check {
		json.NewEncoder(w).Encode(map[string]string{"status": "Error in updating task"})
	}
}

func handleRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", homePage).Methods("Get")
	router.HandleFunc("/gettasks", getAllTasks).Methods("Get")
	router.HandleFunc("/gettask/{id}", task).Methods("Get")
	router.HandleFunc("/create", createTask).Methods("POST")
	router.HandleFunc("/delete/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/update/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8082", router))
}

func main() {
	allTasks()
	fmt.Println("Hello Flutter boys")
	handleRoutes()

}
