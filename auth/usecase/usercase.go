package auth_usecase

import (
	"encoding/json"

	"github.com/genku-m/upsider-cording-test/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

type AuthUsecase struct {
	AuthRepository AuthRepository
}

func NewAuthUsecase(authRepo AuthRepository) *AuthUsecase {
	return &AuthUsecase{
		AuthRepository: authRepo,
	}
}

func (u *AuthUsecase) Login(ctx *gin.Context, email, password string) error {
	user, err := u.AuthRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return err
	} else {
		session := sessions.Default(ctx)
		loginUser, err := json.Marshal(&models.LoginInfo{
			GUID:        user.GUID,
			CompanyGUID: user.CompanyGUID,
		})
		if err == nil {
			session.Set("loginUser", string(loginUser))
			session.Save()
		} else {
			return err
		}
	}

	return nil
}
