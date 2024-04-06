package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type serverConfig struct{}

type server struct {
	config *serverConfig
}

func NewConfig() *serverConfig {
	return &serverConfig{}
}

func NewServer(cfg *serverConfig) *server {
	return &server{
		config: cfg,
	}
}

func (s *server) Listen() error {
	router := gin.Default()

	router.POST("/api/invoices", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	router.GET("/api/invoices", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	return router.Run()
}
