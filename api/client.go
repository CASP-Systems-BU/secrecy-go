package api

// A client class that creates connection to the service and call the exported routes.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Host string
}

func NewClient(host string) *Client {
	return &Client{
		Host: host,
	}
}

func (c *Client) sendRequest(requestType string, requestRoute string, requestBody []byte) (string, error) {
	url := fmt.Sprintf("%s/%s", c.Host, requestRoute)
	req, err := http.NewRequest("GET", url, bytes.NewBuffer(requestBody))
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

func (c *Client) PostPing(requestTime string) (string, error) {
	requestBody, err := json.Marshal(map[string]string{
		"request_time": requestTime,
	})
	if err != nil {
		return "", err
	}

	return c.sendRequest("POST", "ping", requestBody)
}

func (c *Client) GetPing(requestTime string) (string, error) {
	return c.sendRequest("GET", "ping", nil)
}
