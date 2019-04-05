package controller

import (
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/kataras/iris"
	"iissy.com/access"
	"iissy.com/models"
	"iissy.com/utils"
)

// Index 首页
func Index(ctx iris.Context) {
	result, err := access.Index()
	if err != nil {
		log.Fatal(err)
	}
	ctx.ViewData("title", "Home page")
	ctx.ViewData("body", result)
	ctx.View("index.html")
}

// Detail 详情
func Detail(ctx iris.Context) {
	id := ctx.Params().Get("id")

	result, err := access.Detail(id)
	if err != nil {
		log.Fatal(err)
	}

	ctx.ViewData("body", result)
	ctx.View("article/detail.html")
}

// Login 登录
func Login(ctx iris.Context) {
	ctx.View("user/login.html")
}

// Loginpost 登陆
func Loginpost(ctx iris.Context) {
	user := models.User{}
	user.UserID = ctx.FormValue("UID")
	user.Password = utils.GetMd5String(ctx.FormValue("PWD"))
	result, err := access.Login(user)
	if err != nil {
		log.Fatal(err)
	}

	if result.ID > 0 {
		id := strconv.Itoa(result.ID)
		roleid := strconv.Itoa(result.RoleID)
		// expiration := time.Now()
		// expiration = expiration.AddDate(0, 0, 1)
		ctx.SetCookieKV("id", id, iris.CookieEncode(sc.Encode))
		ctx.SetCookieKV("roleid", roleid, iris.CookieEncode(sc.Encode))
		ctx.SetCookieKV("userid", result.UserID, iris.CookieEncode(sc.Encode))
		ctx.SetCookieKV("username", url.QueryEscape(result.UserName), iris.CookieEncode(sc.Encode))
		ctx.SetCookieKV("token", utils.Encryption(id, roleid, result.UserID), iris.CookieEncode(sc.Encode))
	}

	var msg models.Uploador
	msg.Success = result.ID > 0
	ctx.JSON(msg)
}

// Logout 注销
func Logout(ctx iris.Context) {
	ctx.RemoveCookie("id")
	ctx.RemoveCookie("userid")
	ctx.RemoveCookie("username")
	ctx.RemoveCookie("token")

	ctx.Redirect("/")
}

// Reg 注册
func Reg(ctx iris.Context) {
	ctx.View("user/reg.html")
}

// Regpost 注册
func Regpost(ctx iris.Context) {
	var user models.User
	var msg models.Uploador
	user.UserID = ctx.FormValue("UserId")
	user.UserName = ctx.FormValue("UserName")
	user.Password = ctx.FormValue("Password")
	if strings.TrimSpace(user.UserID) == "" || strings.TrimSpace(user.UserName) == "" || strings.TrimSpace(user.Password) == "" {
		msg.Success = false
	} else {
		user.Password = utils.GetMd5String(user.Password)
		result, err := access.RegPost(user)
		utils.CheckErr(err)
		msg.Success = result
	}

	ctx.JSON(msg)
}
