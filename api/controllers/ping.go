package controllers

import (
	"net/http"
	"time"

	"github.com/CASP-Systems-BU/secrecy-go/api/models"
	"github.com/gin-gonic/gin"
)

// PingGet is a function that handles GET requests to /ping.
// It returns a JSON response with the request time(seen by server) and the response time.
func PingGet(c *gin.Context) {
	var pingRequest models.PingRequest

	pingRequest.RequestTime = c.Query("request_time")
	pingRequest.ResponseTime = time.Now().Format(time.RFC3339)

	c.JSON(http.StatusOK, pingRequest)
}

// PingPost is a function that handles POST requests to /ping.
// It returns a JSON response with the request time(seen by client) and the response time.
func PingPost(c *gin.Context) {
	var pingRequest models.PingRequest

	c.BindJSON(&pingRequest)
	pingRequest.ResponseTime = time.Now().Format(time.RFC3339)

	c.JSON(http.StatusOK, pingRequest)
}
