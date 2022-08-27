package routers

import (
	"github.com/gin-gonic/gin"
	v1 "goBlog/api/v1"
	"goBlog/midderware"
)

func InitRoute() {
	gin.SetMode("debug")
	r := gin.New()
	r.Use(midderware.Logger())
	r.Use(midderware.Cors())
	r.Use(gin.Recovery())

	auth := r.Group("api/v1")
	auth.Use(midderware.JwtToken())
	{
		//User模块路由

		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		//Article模块路由
		auth.POST("article/add", v1.AddArticle)
		auth.PUT("article/:id", v1.EditArticle)
		auth.DELETE("article/:id", v1.DeleteArticle)

		//Category模块路由
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCategory)
		auth.DELETE("category/:id", v1.DeleteCategory)
		//上传文件
		auth.POST("upload", v1.UpLoad)
		//获得个人信息
		auth.PUT("profile/:id", v1.UpdateProfile)
	}
	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("users", v1.GetUsers)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("article", v1.GetArticle)
		router.GET("article/list/:id", v1.GetAllArticle)
		router.GET("article/:id", v1.GetArticleInfo)
		router.GET("category", v1.GetCategory)
		router.GET("category/:id", v1.GetCategoryInfo)
		router.POST("login", v1.Login)
		//获得个人信息
		router.GET("profile/:id", v1.GetProfile)
	}
	r.Run(":3000")

}
