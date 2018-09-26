package controllers

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/julienschmidt/httprouter"
	"iissy.com/ado"
	"iissy.com/models"
	"iissy.com/utils"
)

// Login is yes
func Login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles(
		"public/views/user/login.html",
		"public/views/shared/_header.html",
		"public/views/shared/_toper.html",
		"public/views/shared/_footer.html")
	err := temp.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// LoginPost is yes
func LoginPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	utils.CheckErr(err)
	user := models.User{}
	user.UserID = r.PostForm["UID"][0]
	user.Password = utils.GetMd5String(r.PostForm["PWD"][0])
	result, err := ado.Login(user)
	if err != nil {
		log.Fatal(err)
	}

	if result.ID > 0 {
		idString := strconv.Itoa(result.ID)
		expiration := time.Now()
		expiration = expiration.AddDate(0, 0, 1)
		idCookie := http.Cookie{Name: "id", Value: idString, Expires: expiration}
		uidCookie := http.Cookie{Name: "userid", Value: result.UserID, Expires: expiration}
		nameCookie := http.Cookie{Name: "username", Value: url.QueryEscape(result.UserName), Expires: expiration}
		tokenCookie := http.Cookie{Name: "token", Value: utils.Encryption(idString, result.UserID), Expires: expiration}
		http.SetCookie(w, &idCookie)
		http.SetCookie(w, &uidCookie)
		http.SetCookie(w, &nameCookie)
		http.SetCookie(w, &tokenCookie)
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
	tokenCookie := http.Cookie{Name: "token", Value: "", Expires: expiration}
	http.SetCookie(w, &idCookie)
	http.SetCookie(w, &uidCookie)
	http.SetCookie(w, &nameCookie)
	http.SetCookie(w, &tokenCookie)

	var msg models.Uploador
	msg.Success = true
	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}

// Reg is yes
func Reg(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles(
		"public/views/user/reg.html",
		"public/views/shared/_header.html",
		"public/views/shared/_toper.html",
		"public/views/shared/_footer.html")
	err := temp.Execute(w, nil)
	if err != nil {
		fmt.Fprintf(w, "%q", err)
	}
}

// RegPost is yes
func RegPost(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	err := r.ParseForm()
	utils.CheckErr(err)

	var user models.User
	var msg models.Uploador
	user.UserID = r.PostForm["UserId"][0]
	user.UserName = r.PostForm["UserName"][0]
	user.Password = r.PostForm["Password"][0]
	if strings.TrimSpace(user.UserID) == "" || strings.TrimSpace(user.UserName) == "" || strings.TrimSpace(user.Password) == "" {
		msg.Success = false
	} else {
		user.Password = utils.GetMd5String(user.Password)
		result, err := ado.RegPost(user)
		utils.CheckErr(err)
		msg.Success = result
	}

	b, _ := json.Marshal(msg)
	fmt.Fprintf(w, "%s", string(b))
}

// UserIndex is yes
func UserIndex(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	temp, _ := template.ParseFiles(
		"public/views/user/user.html",
		"public/views/shared/_menu.html",
		"public/views/shared/_header.html")
	id, _ := utils.GetUser(r)

	result, err := ado.Get(id)
	utils.CheckErr(err)

	result.UserName += " - 个人中心"
	err = temp.Execute(w, result)
	utils.CheckErr(err)
}
