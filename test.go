package main

import (
	"log"
	"net/url"
)

func main() {
	//user := "t2_11qnzrqvOKJO"

	//validUserName := regexp.MustCompile("t2_+[a-z-0-9_]*$")
	//fmt.Println(validUserName.MatchString(user))

	u, err := url.ParseRequestURI("hi/there?")
	log.Printf("hi/there?: err=%+v url=%+v\n", err, u)

	u, err = url.ParseRequestURI("http://golang.cafe/")
	log.Printf("hi/there?: err=%+v url=%+v\n", err, u)

	u, err = url.ParseRequestURI("http://golang.org/index.html?#page1")
	log.Printf("hi/there?: err=%+v url=%+v\n", err, u)

	u, err = url.ParseRequestURI("golang.org")
	log.Printf("hi/there?: err=%+v url=%+v\n", err, u)
}
