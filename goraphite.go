package goraphite

import "fmt"

// Client is the main client for talking to Graphite API
type Client struct {
	Host string
	Port int
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
