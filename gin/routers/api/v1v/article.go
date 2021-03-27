package v1v

import (
	"github.com/Michael-Johy/go-pkgs-practise/gin/models/blog"
	"github.com/Michael-Johy/go-pkgs-practise/gin/pkg/e"
	"github.com/Michael-Johy/go-pkgs-practise/gin/pkg/setting"
	"github.com/Michael-Johy/go-pkgs-practise/gin/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"net/http"
)

func GetArticles(c *gin.Context) {
	maps := make(map[string]interface{})
	title := c.Query("title")
	if title != "" {
		maps["title"] = title
	}
	datas := blog.GetArticle(util.GetPage(c), setting.PageSize, maps)

	data := make(map[string]interface{})
	data["list"] = datas
	data["total"] = blog.GetArticleCount(maps)

	c.JSON(http.StatusOK, util.BuildRes(e.SUCCESS, data))
}

func CreateArticle(c *gin.Context) {
	var article blog.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if blog.CreateArticle(article) {
		c.JSON(http.StatusOK, util.BuildRes(e.SUCCESS, make(map[string]interface{})))
	} else {
		c.JSON(http.StatusBadRequest, util.BuildRes(e.BAD_REQUEST, make(map[string]interface{})))
	}
}

func UpdateArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	var article blog.Article
	if err := c.ShouldBindJSON(&article); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if blog.UpdateArticle(id, article) {
		c.JSON(http.StatusOK, util.BuildRes(e.SUCCESS, make(map[string]interface{})))
	} else {
		c.JSON(http.StatusBadRequest, util.BuildRes(e.BAD_REQUEST, make(map[string]interface{})))

	}
}

func DeleteArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	if blog.DeleteArticle(id) {
		c.JSON(http.StatusOK, util.BuildRes(e.SUCCESS, make(map[string]interface{})))
	} else {
		c.JSON(http.StatusBadRequest, util.BuildRes(e.BAD_REQUEST, make(map[string]interface{})))

	}
}
