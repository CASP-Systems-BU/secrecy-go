package models

// PingRequest is a struct that represents the request and response time of a ping request.
// It is used to marshal and unmarshal JSON for the /ping route.
type PingRequest struct {
	RequestTime  string `json:"request_time"`
	ResponseTime string `json:"response_time"`
}
