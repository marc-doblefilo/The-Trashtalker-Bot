package main

import (
	"fmt"
	"os"
	"time"
	"encoding/json"
	"net/http"
	"github.com/joho/godotenv"
)

var myClient = &http.Client{Timeout: 10 * time.Second}


type Board struct {
	Name string `json:"name"`
	Id string	`json:"id"`
	Url string	`json:"url"`
}

func main() {

	godotenv.Load(".env.dist")

	Secret := os.Getenv("TrelloSecret")
	API := os.Getenv("TrelloAPI")
	url := fmt.Sprintf("https://api.trello.com/1/members/me/boards?key=%s&token=%s", API, Secret)


	r, err := myClient.Get(url)
    if err != nil {
        fmt.Println(err)
    }
	defer r.Body.Close()
	boars := []Board{}
	
	json.NewDecoder(r.Body).Decode(&boars)
	
	fmt.Println(boars[0].Name)
	fmt.Println(boars)
}
