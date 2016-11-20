package main

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	example "github.com/elauqsap/go-api-pkg-example"
)

func main() {
	proxy, _ := url.Parse("http://localhost:8080")
	c := example.NewClient(&http.Client{Transport: &http.Transport{Proxy: http.ProxyURL(proxy)}}, nil)
	// create a user
	u := example.UserBody{
		Login:    "example@test.com",
		Password: "example",
	}
	data, _ := json.Marshal(u)
	resp, err := c.User.Create(data)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	} else {
		log.Printf("%#v", resp)
	}
	// return the user from the db
	resp, err = c.User.Read(1)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	} else {
		log.Printf("%#v", resp)
	}
	// update the user in the db
	u.Password = "updated"
	data, _ = json.Marshal(u)
	resp, err = c.User.Update(1, data)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	} else {
		log.Printf("%#v", resp)
	}
	// delete the user from the db
	resp, err = c.User.Delete(1)
	if err != nil {
		log.Printf("Error: %s", err.Error())
	} else {
		log.Printf("%#v", resp)
	}
}
