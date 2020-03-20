package api

import (
	"github.com/gin-gonic/gin"
	"hrefs.cn/src/controller"
	"hrefs.cn/src/middleware"
)

func Start() {
	app := gin.New()
	app.Use(gin.Recovery())
	app.Use(middleware.TraceApi)
	app.POST("/login", controller.Login)
	api := app.Group("/api")
	api.Use(middleware.CheckLogin)
	api.POST("/upload", controller.Upload)

	article := api.Group("/article")
	{
		article.POST("/save", controller.SaveArticle)
		article.POST("/list/:size/:page", controller.GetArticleList)
		article.GET("/get/:id", controller.GetArticle)
		article.GET("/delete/:id", controller.DeleteArticle)
	}

	account := api.Group("/account")
	{
		account.POST("/list/:size/:page", controller.GetAccountList)
		account.GET("/get/:id", controller.GetAccount)
	}

	link := api.Group("/link")
	{
		link.POST("/save", controller.SaveLink)
		link.POST("/list/:size/:page", controller.GetLinkList)
		link.GET("/get/:id", controller.GetLink)
		link.GET("/cat/list", controller.GetCatOptions)
		link.GET("/delete/:id", controller.DeleteLink)
	}

	cuslink := api.Group("/cuslink")
	{
		cuslink.POST("/save", controller.SaveCusLink)
		cuslink.POST("/list/:size/:page", controller.GetCusLinkList)
		cuslink.GET("/get/:id", controller.GetCusLink)
		cuslink.GET("/delete/:id", controller.DeleteCusLink)
	}

	app.Run(":81")
}
