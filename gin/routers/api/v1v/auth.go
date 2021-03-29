package v1v

import (
	"github.com/Michael-Johy/go-pkgs-practise/gin/models/blog"
	"github.com/Michael-Johy/go-pkgs-practise/gin/pkg/e"
	"github.com/Michael-Johy/go-pkgs-practise/gin/pkg/util"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthVO struct {
	Username string `valid:"Required; MaxSize(50)" json:"username"`
	Password string `valid:"Required; MaxSize(50)" json:"password"`
}

func Login(c *gin.Context) {
	var authVo AuthVO
	if err := c.ShouldBindJSON(&authVo); err != nil {
		c.JSON(http.StatusBadRequest, util.BuildRes(e.BAD_REQUEST, make(map[string]interface{})))
	} else {
		valid := validation.Validation{}
		isValid, _ := valid.Valid(authVo)
		if !isValid || !blog.ExistUser(authVo.Username, authVo.Password) {
			c.JSON(http.StatusBadRequest, util.BuildRes(e.BAD_REQUEST, make(map[string]interface{})))
		} else {
			token, err := util.GenerateToken(authVo.Username, authVo.Password)
			if err != nil {
				c.JSON(http.StatusBadRequest, util.BuildRes(e.INVALID_PARAM, make(map[string]interface{})))
			} else {
				data := make(map[string]interface{})
				data["token"] = token
				c.JSON(http.StatusOK, util.BuildRes(e.SUCCESS, data))
			}
		}

	}
}
