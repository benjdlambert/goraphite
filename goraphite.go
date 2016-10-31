package goraphite

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/benjdlambert/goraphite/models"
	"github.com/benjdlambert/goraphite/query"
)

// Client is the main client for talking to Graphite API
type Client struct {
	Host string
	Port int
}

// Status checks the current health of the Graphite API
func (c *Client) Status() (*Status, error) {
	response, err := c.request("/")

	if err != nil {
		return nil, err
	}

	if response.StatusCode >= 400 {
		return nil, fmt.Errorf(
			"Response code: %d recieved from Graphite API",
			response.StatusCode,
		)
	}

	return &Status{response.StatusCode}, nil
}

// FindMetrics uses the /metrics/find endpoint to list metrics
func (c *Client) FindMetrics(query query.FindOptions) (*[]models.Metric, error) {
	target := []models.Metric{}
	queryString, err := query.String()

	if err != nil {
		return nil, err
	}

	path := fmt.Sprintf("/metrics/find?%s", queryString)

	err = c.jsonRequest(path, &target)

	if err != nil {
		return nil, err
	}

	return &target, nil
}

func (c *Client) jsonRequest(path string, target interface{}) error {
	response, err := c.request(path)

	if err != nil {
		return err
	}

	if response.StatusCode != 200 {
		return fmt.Errorf(
			"Received response code %d for path %s",
			response.StatusCode,
			path,
		)
	}

	defer response.Body.Close()

	return json.NewDecoder(response.Body).Decode(target)
}

func (c *Client) request(path string) (*http.Response, error) {
	return http.Get(
		fmt.Sprintf("http://%s:%d%s", c.Host, c.Port, path),
	)
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
