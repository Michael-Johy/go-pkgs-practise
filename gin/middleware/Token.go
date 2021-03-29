package middleware

import (
	"github.com/Michael-Johy/go-pkgs-practise/gin/models/blog"
	"github.com/Michael-Johy/go-pkgs-practise/gin/pkg/e"
	"github.com/Michael-Johy/go-pkgs-practise/gin/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var code int
		var data interface{}

		token := ctx.GetHeader("token")
		if token == "" {
			code = e.INVALID_PARAM
		} else {
			claims, err := util.ParseToken(token)
			if err == nil && claims != nil {
				username := claims.Username
				password := claims.Password
				if blog.ExistUser(username, password) {
					code = e.SUCCESS
				}
			} else {
				code = e.INVALID_PARAM
			}
		}

		if code != e.SUCCESS {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"code":    code,
				"message": e.GetMsg(code),
				"data":    data,
			})
			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
