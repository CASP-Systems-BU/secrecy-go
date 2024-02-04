package controllers

import (
	"net/http"
	"time"

	"github.com/CASP-Systems-BU/secrecy-go/api/models"
	"github.com/gin-gonic/gin"
)

func PingGet(c *gin.Context) {
	var pingRequest models.PingRequest

	pingRequest.RequestTime = c.Query("request_time")
	pingRequest.ResponseTime = time.Now().Format(time.RFC3339)

	c.JSON(http.StatusOK, pingRequest)
}

func PingPost(c *gin.Context) {
	var pingRequest models.PingRequest

	c.BindJSON(&pingRequest)
	pingRequest.ResponseTime = time.Now().Format(time.RFC3339)

	c.JSON(http.StatusOK, pingRequest)
}
