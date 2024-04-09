package auth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/genku-m/upsider-cording-test/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
)

func LoginCheckMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		session := sessions.Default(ctx)
		loginUserJson, err := dproxy.New(session.Get("loginUser")).String()
		fmt.Println("test")
		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			ctx.Abort()
		} else {
			var loginInfo models.LoginInfo
			err := json.Unmarshal([]byte(loginUserJson), &loginInfo)
			if err != nil {
				ctx.Status(http.StatusUnauthorized)
				ctx.Abort()
			} else {
				ctx.Next()
			}
		}
	}
}
