package board

import (
	"fmt"
	"os"
	"encoding/json"
	"net/http"
	"github.com/joho/godotenv"
)

type Board struct {
	Name string `json:"name"`
	Id string	`json:"id"`
	Url string	`json:"url"`
}



func getBoards(boards []Board) []Board {
	var myClient = &http.Client{}
		
	godotenv.Load(".env.dist")

	Secret := os.Getenv("TrelloSecret")
	API := os.Getenv("TrelloAPI")
	url := fmt.Sprintf("https://api.trello.com/1/members/me/boards?key=%s&token=%s", API, Secret)

	r, err := myClient.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&boards)

	return boards
}