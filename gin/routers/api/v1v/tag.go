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

func GetTags(c *gin.Context) {
	maps := make(map[string]interface{})
	name := c.Query("name")
	if name != "" {
		maps["name"] = name
	}
	state := c.Query("state")
	if state != "" {
		maps["state"] = com.StrTo(state).MustInt()
	}

	data := make(map[string]interface{})
	data["list"] = blog.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = blog.GetTagCount(maps)

	c.JSON(http.StatusOK, util.BuildRes(e.SUCCESS, data))
}

func CreateTag(c *gin.Context) {
	var tag blog.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if blog.CreateTag(tag) {
		c.JSON(http.StatusOK, util.BuildRes(e.SUCCESS, make(map[string]interface{})))
	} else {
		c.JSON(http.StatusBadRequest, util.BuildRes(e.BAD_REQUEST, make(map[string]interface{})))
	}
}

func UpdateTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	var tag blog.Tag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if blog.UpdateTag(id, tag) {
		c.JSON(http.StatusOK, util.BuildRes(e.SUCCESS, make(map[string]interface{})))
	} else {
		c.JSON(http.StatusBadRequest, util.BuildRes(e.BAD_REQUEST, make(map[string]interface{})))

	}
}

func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt64()
	if blog.DeleteTags(id) {
		c.JSON(http.StatusOK, util.BuildRes(e.SUCCESS, make(map[string]interface{})))
	} else {
		c.JSON(http.StatusBadRequest, util.BuildRes(e.BAD_REQUEST, make(map[string]interface{})))

	}
}
