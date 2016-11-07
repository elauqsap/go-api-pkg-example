package client

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// User ...
type User struct {
	Client *Client
}

// UserBody ...
type UserBody struct {
	ID    int    `json:"id,omitempty"`
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Role  string `json:"role,omitempty"`
	Key   string `json:"api_key,omitempty"`
}

// Service ...
func (s *User) Service(u UserBody, method string) (*Response, error) {
	if !s.Client.Methods[method] {
		return nil, fmt.Errorf("%s is not a valid HTTP method for the user api", method)
	}
	data, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(data)
	resp := new(Response)
	if err = s.Client.Request(method, "user", body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
