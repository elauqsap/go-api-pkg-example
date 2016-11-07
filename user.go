package client

import (
	"bytes"
	"encoding/json"
)

// UserService ...
type UserService struct {
	Client *Client
}

// User ...
type User struct {
	ID    int    `json:"id,omitempty"`
	First string `json:"first"`
	Last  string `json:"last"`
	Role  string `json:"role"`
	Key   string `json:"api_key,omitempty"`
}

// Create ...
func (s *UserService) Create(u User) (*Response, error) {
	data, err := json.Marshal(u)
	if err != nil {
		return nil, err
	}
	body := bytes.NewReader(data)
	resp := new(Response)
	if err = s.Client.Post("user", body, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
