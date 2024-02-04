package models

// Ping is a model for the ping route.
type PingRequest struct {
	RequestTime  string `json:"request_time"`
	ResponseTime string `json:"response_time"`
}
