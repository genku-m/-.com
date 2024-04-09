package auth

import (
	"encoding/json"

	"github.com/genku-m/upsider-cording-test/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
)

func LoginCheck(ctx *gin.Context) (*models.LoginInfo, error) {
	session := sessions.Default(ctx)
	loginUserJson, err := dproxy.New(session.Get("loginUser")).String()
	if err != nil {
		return nil, err
	}
	var loginInfo models.LoginInfo
	err = json.Unmarshal([]byte(loginUserJson), &loginInfo)
	if err != nil {
		return nil, err
	}

	return &loginInfo, nil
}
