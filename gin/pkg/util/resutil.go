package util

import (
	"github.com/Michael-Johy/go-pkgs-practise/gin/pkg/e"
	"github.com/gin-gonic/gin"
)

func BuildRes(code int, data map[string]interface{}) gin.H {
	return gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	}
}
