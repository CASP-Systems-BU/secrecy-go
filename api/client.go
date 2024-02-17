package api

// A client class that creates connection to the service and call the exported routes.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Client is a struct that manages requests to
// a secrecy service on one of the computing parties.
type Client struct {
	Host string
}

// NewClient is a function that returns a new Client for the specified host.
func NewClient(host string) *Client {
	return &Client{
		Host: host,
	}
}

// sendRequest is a function that sends a request to the service.
// It is used by the other methods to send requests to the service after getting
// the marshled body.
// It returns the response body and an error if any.
func (c *Client) sendRequest(requestType string, requestRoute string, requestBody []byte) (string, error) {
	url := fmt.Sprintf("%s/%s", c.Host, requestRoute)
	req, err := http.NewRequest(requestType, url, bytes.NewBuffer(requestBody))
	if err != nil {
		return "", err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

// PostPing is a function that sends a POST request to the service.
// It sends a request time and returns the response time.
func (c *Client) PostPing(requestTime string) (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"request_time": requestTime,
	})
	if err != nil {
		return "", err
	}

	return c.sendRequest("POST", "ping", requestBody)
}

// GetPing is a function that sends a GET request to
// the service to get the response time.
func (c *Client) GetPing(requestTime string) (string, error) {
	return c.sendRequest("GET", "ping", nil)
}
