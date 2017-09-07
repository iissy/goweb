package controllers

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
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
	temp, _ := template.New("_list.html").Funcs(template.FuncMap{"daysAgo": daysAgo}).ParseFiles("public/views/index.html", "public/views/_header.html", "public/views/_toper.html", "public/views/_list.html", "public/views/_footer.html")
	result, err := services.Index(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, result)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

func daysAgo(str string) string {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, str)
	day := t.Format("2006-01-02")
	return day
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
	r.ParseForm()
	var user models.User
	user.UserID = r.PostForm["UID"][0]
	user.Password = r.PostForm["PWD"][0]
	result, err := services.Login(user)
	if err != nil {
		log.Fatal(err)
	}

	if result.ID > 0 {
		expiration := time.Now()
		expiration = expiration.AddDate(0, 0, 1)
		idCookie := http.Cookie{Name: "id", Value: strconv.Itoa(result.ID), Expires: expiration}
		uidCookie := http.Cookie{Name: "userid", Value: result.UserID, Expires: expiration}
		nameCookie := http.Cookie{Name: "username", Value: url.QueryEscape(result.UserName), Expires: expiration}
		http.SetCookie(w, &idCookie)
		http.SetCookie(w, &uidCookie)
		http.SetCookie(w, &nameCookie)
	}

	var msg models.Uploador
	msg.Success = result.ID > 0
	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}

// Logout is yes
func Logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	expiration := time.Now()
	expiration = expiration.AddDate(0, 0, -1)
	idCookie := http.Cookie{Name: "id", Value: "", Expires: expiration}
	uidCookie := http.Cookie{Name: "userid", Value: "", Expires: expiration}
	nameCookie := http.Cookie{Name: "username", Value: "", Expires: expiration}
	http.SetCookie(w, &idCookie)
	http.SetCookie(w, &uidCookie)
	http.SetCookie(w, &nameCookie)

	var msg models.Uploador
	msg.Success = true
	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}

// Add is yes
func Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles("public/views/add.html", "public/views/_header.html", "public/views/_toper.html", "public/views/_footer.html")
	err := temp.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// Reg is yes
func Reg(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles("public/views/reg.html", "public/views/_header.html", "public/views/_toper.html", "public/views/_footer.html")
	err := temp.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// RegPost is yes
func RegPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	r.ParseForm()
	var user models.User
	user.UserID = r.PostForm["UserId"][0]
	user.UserName = r.PostForm["UserName"][0]
	user.Password = r.PostForm["Password"][0]
	result, err := services.RegPost(user)
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
	temp, _ := template.ParseFiles("public/views/user.html", "public/views/_header.html", "public/views/_toper.html", "public/views/_footer.html")
	result, err := services.User(ps.ByName("id"))
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, result)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// Search is yes
func Search(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var msg models.Uploador
	msg.Message = "asdf"
	msg.Path = "/images/star.png"
	msg.Success = true
	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}
