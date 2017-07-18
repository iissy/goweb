package main

import (
	"net/http"

	"asyons.com/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("public"))
	router.GET("/", controllers.Index)
	router.GET("/item/:id", basicAuth(controllers.Detail))
	router.GET("/login", controllers.Login)
	router.GET("/logout", controllers.Logout)

	router.GET("/add", controllers.Add)
	router.GET("/list", controllers.List)
	router.POST("/upload", controllers.Upload)
	router.POST("/post", controllers.Post)

	router.GET("/user/:id", controllers.User)
	router.GET("/mine", controllers.Mine)

	router.POST("/search", controllers.Search)

	http.ListenAndServe(":8000", router)
}

func basicAuth(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		if ok := check(r); ok {
			h(w, r, ps)
		} else {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}

func check(r *http.Request) bool {
	cookie, _ := r.Cookie("username")
	if cookie == nil {
		return false
	}
	return true
}
