package api

import (
	"github.com/CASP-Systems-BU/secrecy-go/api/controllers"
	"github.com/gin-gonic/gin"
)

// Service is a struct that contains the router.
type Service struct {
	Router *gin.Engine
}

// NewService is a function that returns a new Service.
func NewService() *Service {
	return &Service{
		Router: gin.Default(),
	}
}

// InitializeRoutes is a function that initializes the routes for the service.
func (s *Service) InitializeRoutes() {
	s.Router.GET("/ping", controllers.PingGet)
	s.Router.POST("/ping", controllers.PingPost)
}

// Run is a function that starts the service.
func (s *Service) Run() {
	s.Router.Run()
}
