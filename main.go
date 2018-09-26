package main

import (
	"net/http"

	"iissy.com/controllers"
	"iissy.com/utils"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("public"))
	router.GET("/", controllers.TechIndex)
	router.GET("/news", controllers.NewsIndex)
	router.POST("/upload", basicAuth(controllers.Upload))

	router.GET("/user/article/:id", basicAuth(controllers.UserArticle))
	router.GET("/course/:id", controllers.Detail)
	router.GET("/article/add", basicAuth(controllers.Add))
	router.GET("/article/get/:id", controllers.GetArticle)
	router.POST("/article/post", basicAuth(controllers.Post))
	router.GET("/article/update/:id", basicAuth(controllers.Add))

	router.GET("/logout", controllers.Logout)
	router.GET("/reg", controllers.Reg)
	router.POST("/regpost", controllers.RegPost)
	router.GET("/login", controllers.Login)
	router.POST("/loginpost", controllers.LoginPost)
	router.GET("/user/index", basicAuth(controllers.UserIndex))

	router.POST("/comment/post", basicAuth(controllers.CommentPost))

	http.ListenAndServe(":80", router)
}

func basicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if ok := utils.Check(r); ok {
			h(w, r, ps)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
