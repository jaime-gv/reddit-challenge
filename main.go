package main

type post struct {
	ID     int    `json:"ID"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

type allPost []post

var tasks = allPost{
	{
		ID:     1,
		Title:  "My first Post",
		Author: "Jaime",
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
}
