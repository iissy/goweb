package main

import (
	"github.com/kataras/iris"
	"iissy.com/controller"
	"iissy.com/utils"
)

func main() {
	app := iris.New()
	app.StaticServe("./public", "/")
	tmpl := iris.HTML("./views", ".html")
	tmpl.Layout("shared/layout.html")
	tmpl.Reload(true)
	tmpl.AddFunc("daysString", utils.DaysString)
	tmpl.AddFunc("pageChanging", utils.PageChanging)
	app.RegisterView(tmpl)

	app.UseGlobal(controller.Before)

	app.Get("/", controller.Index)
	app.Get("/course/{id}", controller.Detail)

	app.Get("/logout", controller.Logout)
	app.Post("/account/list/{size}/{page}", controller.BasicAuth(controller.Accountlist))

	app.Post("/upload", controller.BasicAuth(controller.Upload))

	app.Get("/login", controller.Webpack)
	app.Get("/reg", controller.Webpack)
	app.Post("/loginpost", controller.Loginpost)
	app.Post("/regpost", controller.Regpost)

	app.Get("/main/{action:path}", controller.BasicAuth(controller.Webpack))

	article := app.Party("article")
	{
		article.Post("/post", controller.BasicAuth(controller.Postarticle))
		article.Post("/list/{size}/{page}", controller.BasicAuth(controller.Articlelist))
		article.Get("/get/{id}", controller.BasicAuth(controller.Getarticle))
		article.Get("/delete/{id}", controller.BasicAuth(controller.Delarticle))
	}

	role := app.Party("role")
	{
		role.Post("/post", controller.BasicAuth(controller.Postrole))
		role.Post("/list/{size}/{page}", controller.BasicAuth(controller.Rolelist))
		role.Get("/get/{id}", controller.BasicAuth(controller.Getrole))
	}

	function := app.Party("function")
	{
		function.Post("/post", controller.BasicAuth(controller.Postfunction))
		function.Post("/list/{size}/{page}", controller.BasicAuth(controller.Functionlist))
		function.Get("/get/{id}", controller.BasicAuth(controller.Getfunction))
		function.Get("/group/{id}", controller.BasicAuth(controller.Functiongroup))
		function.Post("/mapping/post", controller.BasicAuth(controller.Mappingpost))
	}

	app.Run(iris.Addr(":80"))
}
