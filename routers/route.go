package routers

import (
	"github.com/gin-gonic/gin"
	"goBlog/utils"
	"net/http"
)

func InitRoute() {
	gin.SetMode("debug")
	router := gin.Default()

	v1Group := router.Group("api/v1")
	{
		v1Group.GET("hello", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"msg": "ok",
			})
		})
	}

	router.Run(utils.HttpPort)

}
