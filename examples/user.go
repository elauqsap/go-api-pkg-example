package main

import (
	"log"
	"net/http"
	"net/url"

	example "github.com/elauqsap/go-api-pkg-example"
)

func main() {
	proxy, _ := url.Parse("http://localhost:8080")
	c := example.NewClient(&http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}, nil)
	u := example.UserBody{
		First: "Test",
		Last:  "User",
		Role:  "user",
	}
	resp, err := c.User.Service(u, http.MethodPost)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	} else {
		log.Printf("%v", resp)
	}
	u = example.UserBody{
		ID: 1,
	}

	// FIXME: issue comes from the echo router not allowing a body in the GET method
	resp, err = c.User.Service(u, http.MethodGet)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	} else {
		log.Printf("%v", resp)
	}
}
