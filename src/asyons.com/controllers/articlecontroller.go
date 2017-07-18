package controllers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"encoding/json"

	"asyons.com/models"
	"asyons.com/services"
	"github.com/julienschmidt/httprouter"
)

// Index is yes
func Index(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles("public/views/index.html", "public/views/_header.html", "public/views/_toper.html", "public/views/_list.html", "public/views/_footer.html")
	result, err := services.Index(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, result)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// Detail is yes
func Detail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles("public/views/detail.html", "public/views/_header.html", "public/views/_toper.html", "public/views/_footer.html")
	result, err := services.Detail(ps.ByName("id"))
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, result)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// Login is yes
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, 1)
	cookie := http.Cookie{Name: "username", Value: "jimmy", Expires: expiration}
	http.SetCookie(w, &cookie)
}

// Logout is yes
func Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, -1)
	cookie := http.Cookie{Name: "username", Value: "jimmy", Expires: expiration}
	http.SetCookie(w, &cookie)
}

// Add is yes
func Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles("public/views/add.html", "public/views/_header.html", "public/views/_toper.html", "public/views/_footer.html")
	err := temp.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// List is yes
func List(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles("public/views/list.html", "public/views/_header.html", "public/views/_toper.html", "public/views/_footer.html")
	err := temp.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// Upload is yes
func Upload(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var result models.Uploador
	result.Message = "asdf"
	result.Path = "/images/star.png"
	result.Success = true
	b, _ := json.Marshal(result)
	fmt.Fprintf(w, "%s", string(b))
}

// Post is yes
func Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var article models.Article
	result, err := services.Post(article)
	if err != nil {
		log.Fatal(err)
	}

	var msg models.Uploador
	msg.Message = "asdf"
	msg.Path = "/images/star.png"
	msg.Success = result
	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}

// User is yes
func User(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var article models.Article
	result, err := services.Post(article)
	if err != nil {
		log.Fatal(err)
	}

	var msg models.Uploador
	msg.Message = "asdf"
	msg.Path = "/images/star.png"
	msg.Success = result
	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}

// Mine is yes
func Mine(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var article models.Article
	result, err := services.Post(article)
	if err != nil {
		log.Fatal(err)
	}

	var msg models.Uploador
	msg.Message = "asdf"
	msg.Path = "/images/star.png"
	msg.Success = result
	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}
