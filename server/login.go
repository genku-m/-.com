package server

import (
	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email" bson:"email"`
	Password string `json:"password" bson:"password"`
}

func (s *Server) Login(ctx *gin.Context) error {
	var request LoginRequest
	err := ctx.BindJSON(&request)
	if err != nil {
		return err
	}

	err = s.authUsecase.Login(ctx, request.Email, request.Password)
	if err != nil {
		return err
	}

	return nil
}
