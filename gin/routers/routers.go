package routers

import (
	"github.com/Michael-Johy/go-pkgs-practise/gin/middleware"
	"github.com/Michael-Johy/go-pkgs-practise/gin/routers/api/v1v"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	router.POST("/login", v1v.Login)

	apiV1 := router.Group("/api/v1v")
	apiV1.Use(middleware.JWT())

	apiV1.GET("/tags", v1v.GetTags)
	apiV1.POST("/tags", v1v.CreateTag)
	apiV1.PUT("/tags/:id", v1v.UpdateTag)
	apiV1.DELETE("/tags/:id", v1v.DeleteTag)

	apiV1.GET("/articles", v1v.GetArticles)
	apiV1.POST("/articles", v1v.CreateArticle)
	apiV1.PUT("/articles/:id", v1v.UpdateArticle)
	apiV1.DELETE("/articles/:id", v1v.DeleteArticle)

	router.GET("/ping", func(context *gin.Context) {
		name := context.Param("name")
		context.JSON(http.StatusOK, gin.H{
			"message": "pong",
			"name":    name,
		})
	})

	return router
}
