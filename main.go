package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"regexp"

	"github.com/gorilla/mux"
)

type post struct {
	ID     int    `json:"ID"`
	Title  string `json:"Title"`
	Author string `json:"Author"`
}

type allPost []post

var posts = allPost{
	{
		ID:     1,
		Title:  "My first Post",
		Author: "Jaime",
	},
}

func validateAuthor() {
	user := "t2_11qnzrqv"
	validUserName := regexp.MustCompile("t2_+[a-z-0-9_]*$")
	fmt.Println(validUserName.MatchString(user))
}

func getPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(posts)
}
func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "API Go Redit")
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	log.Fatal(http.ListenAndServe(":3000", router))
}
