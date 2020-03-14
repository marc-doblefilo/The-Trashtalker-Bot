package main

import (
	"fmt"
	"os"
	"time"
	"encoding/json"
	"net/http"
	"github.com/joho/godotenv"

	"board"
)

var myClient = &http.Client{Timeout: 10 * time.Second}


func main() {
	
	godotenv.Load(".env.dist")

	Secret := os.Getenv("TrelloSecret")
	API := os.Getenv("TrelloAPI")
	url := fmt.Sprintf("https://api.trello.com/1/members/me/boards?key=%s&token=%s", API, Secret)


	
}


https://developers.trello.com/reference#cards-1