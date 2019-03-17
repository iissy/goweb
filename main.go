package main

import (
	"github.com/kataras/iris"
	"iissy.com/utils"

	_ "github.com/go-sql-driver/mysql"
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

	app.Get("/status", status)

	app.Get("/reg", reg)
	app.Post("/regpost", regpost)
	app.Get("/login", login)
	app.Post("/login", loginpost)
	app.Get("/logout", logout)

	app.Get("/user/index", basicAuth(user))
	app.Post("/upload", basicAuth(upload))

	app.Get("/article/add", basicAuth(addarticle))
	app.Post("/article/post", basicAuth(postarticle))

	app.Get("/article/list/{page}", basicAuth(articlelist))
	app.Get("/article/get/{id}", basicAuth(getarticle))
	app.Get("/article/update/{id}", basicAuth(addarticle))

	app.Run(iris.Addr(":80"))
}
