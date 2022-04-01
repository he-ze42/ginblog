package router

import (
	"fmt"
	"ginblog/controller"

	"github.com/gin-gonic/gin"
)

func Start() {
	e := gin.Default()
    //加载静态文件
    e.Static("/assets", "./assets")

	//加载渲染文件地址，/*代表加载这个地址文件下的所有文件
	e.LoadHTMLGlob("templates/*")
	

	//e.GET("/index", controller.ListUser)
	e.POST("/register", controller.RegisterUsers)

	e.POST("/login", controller.Login)
	e.GET("/", controller.Index)
	//跳转到注册页面
	e.GET("/goregister", controller.GoRegister)
	//跳转到登录页面
	e.GET("/gologin", controller.Gologin)

	//操作博客
	//跳转到博客列表
	e.GET("/gopost",controller.GoAddPost)
	//添加博客
	e.POST("/post",controller.AddPost)
	//跳转到增加博客页面
	e.GET("/post_index",controller.GetPostIndex)
	//博客详细页面
	e.GET("/post_detail",controller.GetPostDetail)

   




	e.Run()
}

