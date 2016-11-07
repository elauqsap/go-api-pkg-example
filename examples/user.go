package main

import (
	"log"

	example "github.com/elauqsap/go-api-pkg-example"
)

func main() {
	c := example.NewClient(nil, nil)
	u := example.User{
		First: "Test",
		Last:  "User",
		Role:  "user",
	}
	resp, err := c.User.Create(u)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	} else {
		log.Printf("%v", resp)
	}
}
