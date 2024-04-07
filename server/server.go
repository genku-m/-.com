package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ServerConfig struct{}

type Server struct {
	invoiceUsecase InvoiceUsecase
	config         *ServerConfig
}

func NewConfig() *ServerConfig {
	return &ServerConfig{}
}

func NewServer(invoiceUsecase InvoiceUsecase, cfg *ServerConfig) *Server {
	return &Server{
		invoiceUsecase: invoiceUsecase,
		config:         cfg,
	}
}

func (s *Server) Listen() error {
	router := gin.Default()
	router.ContextWithFallback = true
	router.POST("/api/invoices", func(ctx *gin.Context) {
		res, err := s.CreateInvoice(ctx)
		if err != nil {
			ctx.String(http.StatusInternalServerError, err.Error())
			return
		}
		ctx.JSON(http.StatusOK, res)
	})

	router.GET("/api/invoices", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	router.GET("/health", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})

	return router.Run()
}
