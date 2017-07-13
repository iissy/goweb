package main

import (
	"fmt"
	"github"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.NotFound = http.FileServer(http.Dir("public"))
	router.GET("/", index)
	router.GET("/issue/:id", basicAuth(detail))
	router.GET("/login", login)
	router.GET("/logout", logout)
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

func index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles("public/views/issue.html", "public/views/_list.html")
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, result)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

func detail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("id"))
}

func login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "username", Value: "jimmy", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, -1)
	cookie := http.Cookie{Name: "username", Value: "jimmy", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func check(r *http.Request) bool {
	cookie, _ := r.Cookie("username")
	if cookie == nil {
		return false
	}
	return true
}
