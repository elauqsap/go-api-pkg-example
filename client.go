package client

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	// Basic client API configuration
	libraryVersion = "0.1"
	defaultBaseURL = "http://127.0.0.1:8081/api/v1/"
	userAgent      = "go-client/" + libraryVersion

	// HTTP Headers
	userAgentHeader = "User-Agent"
)

// A Client manages communication with the API
type Client struct {
	// HTTP client used to communicate with the API
	Client *http.Client

	// BaseURL for API requests
	BaseURL *url.URL

	// UserAgent used when communicating with the API
	UserAgent string

	// Services used for talking to different parts of the API
	User *User
}

// Response models the JSON response from API services
type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
}

// NewClient creates a new client that can use the API services
func NewClient(client *http.Client, baseURL *url.URL) *Client {
	if client == nil {
		client = http.DefaultClient
	}

	if baseURL == nil {
		baseURL, _ = url.Parse(defaultBaseURL)
	}

	c := &Client{Client: client, BaseURL: baseURL, UserAgent: userAgent}
	c.User = &User{c}
	return c
}

// NewRequest creates an HTTP request for the client to make
func (c *Client) NewRequest(method, urlStr string, body io.Reader) (*http.Request, error) {
	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	req, err := http.NewRequest(method, u.String(), body)
	if err != nil {
		return nil, err
	}

	if c.UserAgent != "" {
		req.Header.Add(userAgentHeader, c.UserAgent)
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	return req, nil
}

// Do performs the request and determines the response
func (c *Client) Do(req *http.Request, v interface{}) error {
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}

	defer func() {
		// Drain up to 512 bytes and close the body to let the Transport reuse the connection
		io.CopyN(ioutil.Discard, resp.Body, 512)
		resp.Body.Close()
	}()

	if v != nil {
		switch v := v.(type) {
		case *[]byte:
			*v, err = ioutil.ReadAll(resp.Body)
		default:
			body, _ := ioutil.ReadAll(resp.Body)
			err = json.Unmarshal(body, &v)
		}
	}

	return err
}

// Request performs an HTTP GET request for the calling service using the Client
func (c *Client) Request(method string, rest string, body io.Reader, resp interface{}) error {
	req, err := c.NewRequest(method, c.BaseURL.String()+rest, body)
	if err != nil {
		return err
	}
	if err = c.Do(req, resp); err != nil {
		return err
	}
	return nil
}
