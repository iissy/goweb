package controllers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"

	"encoding/json"

	"asyons.com/models"
	"asyons.com/services"
	"asyons.com/utils"
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
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	ext := utils.ExtensionName(handler.Filename)
	path := "/upload/" + utils.UniqueID() + ext
	f, err := os.OpenFile("./public"+path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	result.Success = true
	result.Path = path
	defer f.Close()
	io.Copy(f, file)

	b, _ := json.Marshal(result)
	fmt.Fprintf(w, "%s", string(b))
}

// Post is yes
func Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	var msg models.Uploador
	msg.Success = false

	var body = r.PostForm["Body"][0]
	var ask models.Ask
	ask.ID = utils.UniqueID()
	ask.NickName = "jimmy"
	ask.UserID = 1
	ask.Subject = r.PostForm["Subject"][0]
	ask.Body = template.HTML(body)

	src := string(body)
	//将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllStringFunc(src, strings.ToLower)

	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	src = re.ReplaceAllString(src, "")

	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	src = re.ReplaceAllString(src, "")

	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	src = re.ReplaceAllString(src, "\n")

	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	src = re.ReplaceAllString(src, "\n")

	ask.Description = utils.Substr2(strings.TrimSpace(src), 0, 150)

	result, err := services.Post(ask)
	if err != nil {
		log.Fatal(err)
	}

	msg.Success = result
	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}

// User is yes
func User(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var article models.Ask
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
	var article models.Ask
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

// Search is yes
func Search(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var article models.Ask
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
