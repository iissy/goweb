package router

import (
	"github.com/kataras/iris"
	"hrefs.cn/src/controller"
	"hrefs.cn/src/middleware"
)

func Start() {
	app := iris.New()
	app.Use(iris.LimitRequestBodySize(1024 * 1024))
	app.StaticWeb("/", "./public")
	tmpl := iris.HTML("./views", ".html")
	tmpl.Layout("shared/layout.html")
	tmpl.Reload(true)
	app.RegisterView(tmpl)
	app.Use(middleware.Trace)

	app.Get("/", controller.Index)
	app.Get("/links/{id}", controller.ListLinks)
	app.Get("/cuslinks", controller.ListCusLinks)
	app.Get("/articles", controller.ListArticles)
	app.Get("/article/{id}", controller.Detail)
	app.Get("/payme", controller.Payme)
	app.Get("/link/{id}", controller.GetLinkUrl)
	app.Get("/cuslink/{id}", controller.GetCusLinkUrl)

	app.Post("/login", controller.Login)

	api := app.Party("api")
	api.Use(middleware.CheckLogin)
	api.Post("/upload", controller.Upload)

	article := api.Party("/article")
	{
		article.Post("/save", controller.SaveArticle)
		article.Post("/list/{size}/{page}", controller.GetArticleList)
		article.Get("/get/{id}", controller.GetArticle)
		article.Get("/delete/{id}", controller.DeleteArticle)
	}

	account := api.Party("/account")
	{
		//account.Post("/post", controller.SaveAccount)
		account.Post("/list/{size}/{page}", controller.GetAccountList)
		account.Get("/get/{id}", controller.GetAccount)
		//account.Get("/delete/{id}", controller.DeleteAccount)
	}

	link := api.Party("/link")
	{
		link.Post("/save", controller.SaveLink)
		link.Post("/list/{size}/{page}", controller.GetLinkList)
		link.Get("/get/{id}", controller.GetLink)
		link.Get("/cat/list", controller.GetCatOptions)
		link.Get("/delete/{id}", controller.DeleteLink)
	}

	cuslink := api.Party("/cuslink")
	{
		cuslink.Post("/save", controller.SaveCusLink)
		cuslink.Post("/list/{size}/{page}", controller.GetCusLinkList)
		cuslink.Get("/get/{id}", controller.GetCusLink)
		cuslink.Get("/delete/{id}", controller.DeleteCusLink)
	}

	app.Run(
		iris.Addr(":80"),
		iris.WithoutBanner,
		iris.WithoutServerError(iris.ErrServerClosed),
	)
}
