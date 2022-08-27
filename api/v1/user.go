package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goBlog/model"
	"goBlog/utils/errmsg"
	"goBlog/utils/validator"
	"net/http"
	"strconv"
)

// UserExist 查询用户是否存在
func UserExist(c *gin.Context) {

}

// AddUser 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	var code int
	var msg string
	//if err := c.ShouldBindJSON(&data); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	//}
	_ = c.ShouldBindJSON(&data)
	msg, code = validator.Validate(&data)
	if code != errmsg.SUCCESS {
		c.JSON(http.StatusOK, gin.H{
			"status":  code,
			"message": msg,
		})

		return
	}
	fmt.Println(data.Username)
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}
	if code == errmsg.ErrorUsernameUsed {
		code = errmsg.ErrorUsernameUsed
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	var id int
	id, _ = strconv.Atoi(c.Param("id"))
	data, code := model.GetSingleUser(id)
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"data":    data,
		"message": errmsg.GetErrMsg(code),
	})
}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pageSize"))
	pageNum, _ := strconv.Atoi(c.Query("pageNum"))
	userName := c.Query("userName")
	data, total := model.GetUsers(userName, pageSize, pageNum)
	c.JSON(http.StatusOK, gin.H{
		"status":  errmsg.SUCCESS,
		"data":    data,
		"total":   total,
		"message": errmsg.GetErrMsg(errmsg.SUCCESS),
	})
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code := model.DeleteUser(id)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// EditUser 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	c.ShouldBindJSON(&data)
	code := model.CheckUpUser(id, data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id, &data)
	}
	if code == errmsg.ErrorUsernameUsed {
		c.Abort() //阻止调用后续函数，与之对应的就是c.next()
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"message": errmsg.GetErrMsg(code),
	})
}

// ChangePassword 改变用户密码
func ChangePassword(c *gin.Context) {

}
