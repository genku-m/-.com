package server

import (
	"net/http"
	"os"

	errpkg "github.com/genku-m/upsider-cording-test/invoice/errors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

type DB struct {
	Name     string
	User     string
	Password string
	Address  string
}
type ServerConfig struct {
	DB *DB
}

type Server struct {
	invoiceUsecase InvoiceUsecase
	authUsecase    AuthUsecase
	config         *ServerConfig
}

func NewConfig() *ServerConfig {
	return &ServerConfig{
		DB: &DB{
			Name:     os.Getenv("DB_NAME"),
			User:     os.Getenv("DB_USER"),
			Password: os.Getenv("DB_PASSWORD"),
			Address:  os.Getenv("DB_ADDRESS"),
		},
	}
}

func NewServer(invoiceUsecase InvoiceUsecase, authUsecase AuthUsecase, cfg *ServerConfig) *Server {
	return &Server{
		invoiceUsecase: invoiceUsecase,
		authUsecase:    authUsecase,
		config:         cfg,
	}
}

func (s *Server) Listen() error {
	router := gin.Default()
	router.ContextWithFallback = true
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("mysession", store))
	router.POST("/login", func(ctx *gin.Context) {
		err := s.Login(ctx)
		if err != nil {
			ctx.String(http.StatusUnauthorized, err.Error())
		}
	})

	router.POST("/api/invoices", func(ctx *gin.Context) {
		res, err := s.CreateInvoice(ctx)
		if err != nil {
			errHundler(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, res)
	})

	router.GET("/api/invoices", func(ctx *gin.Context) {
		res, err := s.ListInvoice(ctx)
		if err != nil {
			errHundler(ctx, err)
			return
		}
		ctx.JSON(http.StatusOK, res)
	})
	router.GET("/health", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "")
	})

	return router.Run()
}

func errHundler(ctx *gin.Context, err error) {
	serverError, ok := err.(*errpkg.ServerError)
	if !ok {
		ctx.String(http.StatusInternalServerError, err.Error())
	}
	switch serverError.ErrCode {
	case errpkg.ErrInvalidArgument:
		ctx.String(http.StatusBadRequest, err.Error())
	case errpkg.ErrNotFound:
		ctx.String(http.StatusNotFound, err.Error())
	case errpkg.ErrInternal:
		ctx.String(http.StatusInternalServerError, err.Error())
	default:
		ctx.String(http.StatusInternalServerError, err.Error())
	}
}
