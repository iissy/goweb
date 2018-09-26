package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"iissy.com/ado"
	"iissy.com/models"
	"iissy.com/utils"
)

// UserArticle is yes
func UserArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp := template.New("index.html").Funcs(template.FuncMap{"daysString": utils.DaysString, "pageChanging": utils.PageChanging})
	temp = template.Must(
		temp.ParseFiles(
			"public/views/article/index.html",
			"public/views/shared/_menu.html",
			"public/views/shared/_header.html",
			"public/views/article/_list.html"))

	id, _ := utils.GetUser(r)
	page, _ := strconv.Atoi(ps.ByName("id"))
	result, err := ado.UserArticle(id, page, 15)
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, result)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// GetArticle is yes
func GetArticle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	result, err := ado.GetArticle(id)
	if err != nil {
		log.Fatal(err)
	}

	b, _ := json.Marshal(result)
	fmt.Fprintf(w, "%s", string(b))
}

// Detail is yes
func Detail(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles(
		"public/views/article/detail.html",
		"public/views/article/_aside.html",
		"public/views/shared/_header.less.html",
		"public/views/shared/_toper.html",
		"public/views/shared/_footer.html")
	id := ps.ByName("id")

	result, err := ado.Detail(id)
	if err != nil {
		log.Fatal(err)
	}

	err = temp.Execute(w, result)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// Add is yes
func Add(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	temp, _ := template.ParseFiles(
		"public/views/article/add.html",
		"public/views/shared/_menu.html",
		"public/views/shared/_header.html")
	err := temp.Execute(w, id)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// Post is yes
func Post(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	utils.CheckErr(err)

	var msg models.Uploador
	msg.Success = false

	var body = r.PostForm["Body"][0]
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

	id, nickname := utils.GetUser(r)
	article := models.Article{
		ID:			r.PostForm["Id"][0],
		NickName:	nickname,
		UserID:		id,
		Subject:	r.PostForm["Subject"][0],
		Picture:	r.PostForm["Picture"][0],
		PostType:	r.PostForm["PostType"][0],
		Origin:		r.PostForm["Origin"][0],
		Description:utils.Substr2(strings.TrimSpace(src), 0, 100),
		Body:		template.HTML(body)}

	adding, _ := strconv.ParseBool(r.PostForm["Adding"][0])
	if adding {
		result, err := ado.Post(article)
		if err != nil {
			log.Fatal(err)
		}
		msg.Success = result
	} else {
		result, err := ado.Update(article)
		if err != nil {
			log.Fatal(err)
		}
		msg.Success = result
	}

	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}
