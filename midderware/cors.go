package midderware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		cors.New(cors.Config{
			AllowOrigins:     []string{"*"}, //表示允许所有域
			AllowMethods:     []string{"*"}, // 该字段是必须的，用来列出浏览器的CORS请求会用到哪些HTTP方法
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length", "Authorization"}, //如果想拿到其他字段，就必须在Access-Control-Expose-Headers里面指定
			AllowCredentials: true,                                        //表示是否发送Cookie
			MaxAge:           12 * time.Hour,
		})
	}
}
