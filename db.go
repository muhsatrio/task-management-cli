package main

import (
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
	ID          int    `json:"id"`
	Rev         int    `json:"rev"`
	Name        string `json:"name"`
	Description string `json:"description"`
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

var baseURL = getURLAPI()
var listTasks []Task

func getURLAPI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	user := os.Getenv("USER_DB")
	password := os.Getenv("PASSWORD_DB")
	fmt.Println(user, password)

	return "http://" + user + ":" + password + "@13.250.43.79:5984"
}

func initTasks() {
	resp, err := http.Get(baseURL + "/efishery_task/_all_docs")

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	bodyBytes, _ := ioutil.ReadAll(resp.Body)

	bodyString := string(bodyBytes)

	fmt.Println(bodyString)

	var responseAPI ResponseAPI
	json.Unmarshal([]byte(bodyString), &responseAPI)
	fmt.Println(responseAPI)

}
