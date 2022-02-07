package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
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

func validateURL() {
	u, err := url.ParseRequestURI("hi/there?")
	u, err = url.ParseRequestURI("http://golang.cafe/")
	log.Printf("hi/there?: err=%+v url=%+v\n", err, u)
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
