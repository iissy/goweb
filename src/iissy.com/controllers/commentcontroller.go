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

// CommentPost is yes
func CommentPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	utils.CheckErr(err)

	var msg models.Uploador
	var body = r.PostForm["Body"][0]
	var comment models.Comment
	id, nickname := utils.GetUser(r)
	comment.UserID = id
	comment.NickName = nickname
	aid, _ := strconv.Atoi(r.PostForm["Id"][0])
	comment.AskID = aid
	comment.Body = template.HTML(body)

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

	result, err := ado.CommentPost(comment)
	if err != nil {
		log.Fatal(err)
	}

	msg.Success = result
	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}
