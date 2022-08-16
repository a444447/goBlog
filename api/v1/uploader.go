package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"goBlog/model"
	"goBlog/utils/errmsg"
	"net/http"
)

func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")

	fileSize := fileHeader.Size
	fmt.Println(fileSize)
	url, code := model.UploaderFile(file, fileSize)

	c.JSON(http.StatusOK, gin.H{
		"status":  code,
		"url":     url,
		"message": errmsg.GetErrMsg(code),
	})
}
