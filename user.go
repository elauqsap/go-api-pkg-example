package client

import (
	"bytes"
	"fmt"
	"net/http"
)

// User embeds a client for making the service REST requests
type User struct {
	Client *Client
}

// UserBody models the input data the user can make via each request
type UserBody struct {
	Login    string `json:"login,omitempty"`
	Password string `json:"password,omitempty"`
}

// Create makes a POST request to /api/v1/user with a JSON body content containing the new user
func (u *User) Create(data []byte) (*Response, error) {
	resp := new(Response)
	if err := u.Client.Request(http.MethodPost, "user", bytes.NewReader(data), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Read makes a GET request to /api/v1/user/:id
func (u *User) Read(id int) (*Response, error) {
	resp := new(Response)
	if err := u.Client.Request(http.MethodGet, fmt.Sprintf("user/%d", id), nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Update makes a PUT request to /api/v1/user/:id with a JSON body content containing the user updates
func (u *User) Update(id int, data []byte) (*Response, error) {
	resp := new(Response)
	if err := u.Client.Request(http.MethodPut, fmt.Sprintf("user/%d", id), bytes.NewReader(data), resp); err != nil {
		return nil, err
	}
	return resp, nil
}

// Delete makes a DELETE request to /api/v1/user/:id, removes the user from the db
func (u *User) Delete(id int) (*Response, error) {
	resp := new(Response)
	if err := u.Client.Request(http.MethodDelete, fmt.Sprintf("user/%d", id), nil, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
