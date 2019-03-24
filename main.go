package main

import (
	"github.com/kataras/iris"
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

	app.UseGlobal(before)

	app.Get("/", index)
	app.Get("/course/{id}", detail)

	app.Get("/logout", logout)
	app.Post("/account/list/{size}/{page}", basicAuth(accountlist))

	app.Post("/upload", basicAuth(upload))

	app.Get("/article/add", basicAuth(addarticle))
	app.Post("/article/post", basicAuth(postarticle))

	app.Get("/login", webpack)
	app.Get("/reg", webpack)
	app.Post("/loginpost", loginpost)
	app.Post("/regpost", regpost)

	app.Get("/main/{action:path}", basicAuth(webpack))
	app.Post("/article/list/{size}/{page}", basicAuth(articlelist))
	app.Get("/article/get/{id}", basicAuth(getarticle))
	app.Get("/article/delete/{id}", basicAuth(delarticle))

	app.Run(iris.Addr(":80"))
}
