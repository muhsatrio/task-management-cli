package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

// Task struct
type Task struct {
	ID          string `json:"_id"`
	Rev         string `json:"_rev"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	Tags        []Tag  `json:"tags"`
	CreatedAt   string `json:"createdAt"`
}

// Docs struct
type Docs struct {
	ID    string `json:"id"`
	Key   string `json:"key"`
	Value struct {
		Rev string `json:"rev"`
	} `json:"value"`
}

// ResponseAPI struct
type ResponseAPI struct {
	TotalRows int    `json:"total_rows"`
	Offset    int    `json:"offset"`
	Rows      []Docs `json:"rows"`
}

// Tag struct
type Tag struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

var baseURL = getURLAPI()
var listTasks []Task

func getURLAPI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")

	return "http://" + user + ":" + password + "@13.250.43.79:5984"
}

func getTask(id string) Task {
	i := 0
	for i < len(listTasks) {
		if listTasks[i].ID == id {
			break
		}
		i++
	}
	if i >= len(listTasks) {
		log.Fatalln("error: not found")
		os.Exit(0)
	}
	return listTasks[i]
}

func filterTask(completed bool) []Task {
	var filteredTask []Task
	for i := 0; i < len(listTasks); i++ {
		if listTasks[i].Completed == completed {
			filteredTask = append(filteredTask, listTasks[i])
		}
	}
	return filteredTask
}

func initTasks() {
	resp, err := http.Get(baseURL + "/efishery_task/_all_docs")

	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)

	var responseAPI ResponseAPI
	json.Unmarshal([]byte(bodyString), &responseAPI)
	listTasks = nil
	for _, doc := range responseAPI.Rows {
		resp, _ = http.Get(baseURL + "/efishery_task/" + doc.ID)
		bodyBytes, _ = ioutil.ReadAll(resp.Body)
		bodyString = string(bodyBytes)
		var taskTemp Task
		json.Unmarshal([]byte(bodyString), &taskTemp)
		listTasks = append(listTasks, taskTemp)
	}
}

func changeTask(id string, rev string, newTask Task) {
	json, _ := json.Marshal(newTask)
	fmt.Println(json)
	_, err := http.NewRequest(http.MethodPut, "/efishery_task/"+id+"?rev="+rev, bytes.NewBuffer(json))

	if err != nil {
		log.Fatalln(err)
		os.Exit(0)
	}
}
