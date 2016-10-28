package goraphite

import (
	"fmt"
	"net/http"
)

// Client is the main client for talking to Graphite API
type Client struct {
	Host string
	Port int
}

func (c *Client) Status() (*Status, error) {
	response, err := c.request("/")

	if err != nil {
		return nil, err
	}

	return &Status{response.StatusCode}, nil
}

func (c *Client) request(path string) (*http.Response, error) {
	return http.Get(path)
}

// NewGoraphite creates a client with the host and port provided
func NewGoraphite(host string, port int) (*Client, error) {
	if host == "" {
		return nil, fmt.Errorf("Host should not be nil or empty")
	}

	if port < 0 {
		return nil, fmt.Errorf("Port should be a valid, positive integer")
	}

	return &Client{
		host,
		port,
	}, nil
}
