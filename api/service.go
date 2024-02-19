package api

import (
	"github.com/CASP-Systems-BU/secrecy-go/api/controllers"
	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
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

	s.Router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"},
		AllowCredentials: true,
	}))

	s.Router.GET("/ping", controllers.PingGet)
	s.Router.POST("/ping", controllers.PingPost)
}

// Run is a function that starts the service.
func (s *Service) Run() {
	s.Router.Run()
}
