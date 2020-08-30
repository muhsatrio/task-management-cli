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
	id          int    `json: "id"`
	rev         int    `json: "rev"`
	name        string `json: "name"`
	description string `json: "description`
	createdAt   string `json: "createdAt"`
}

// ResponseAPI struct
type ResponseAPI struct {
	totalRows int `json: "total_rows"`
	offset    int `json: "offset"`
	rows      struct {
		id    string `json: "id"`
		key   string `json: "key"`
		value struct {
			rev string `json: "rev"`
		} `json: "value"`
	} `json: "rows"`
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
