package v1

import (
	"github.com/gin-gonic/gin"
	"goBlog/model"
	"goBlog/utils/errmsg"
	"net/http"
	"strconv"
)

// GetProfile 获得个人信息
func GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetProfile(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// UpdateProfile 更新个人信息
func UpdateProfile(c *gin.Context) {
	var data model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)
	code := model.UpdateProfile(id, &data)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}
