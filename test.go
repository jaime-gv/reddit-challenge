package main

import (
	"fmt"
	"regexp"
)

func main() {
	user := "t2_11qnzrqvOKJO"

	validUserName := regexp.MustCompile("t2_+[a-z-0-9_]*$")
	fmt.Println(validUserName.MatchString(user))

}
