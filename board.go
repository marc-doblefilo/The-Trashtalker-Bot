package board

type Board struct {
	Name string `json:"name"`
	Id string	`json:"id"`
	Url string	`json:"url"`
}


// get, post, update, delete
 

board1 := Board{}
board.1

func getBoards(boards []Board){
	boards := []Board{}

	r, err := myClient.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer r.Body.Close()

	json.NewDecoder(r.Body).Decode(&boards)

	return boards
}