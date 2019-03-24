package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"iissy.com/access"
	"iissy.com/models"
	"iissy.com/utils"
)

var (
	hashKey  = []byte("the-big-and-secret-fash-key-here")
	blockKey = []byte("lot-secret-of-characters-big-too")
	sc       = securecookie.New(hashKey, blockKey)
)

func before(ctx iris.Context) {
	id, username := utils.GetUser(ctx)
	ctx.ViewData("islogin", id > 0)
	ctx.ViewData("username", username)
	ctx.Next()
}

func index(ctx iris.Context) {
	result, err := access.Index()
	if err != nil {
		log.Fatal(err)
	}
	ctx.ViewData("title", "Home page")
	ctx.ViewData("body", result)
	ctx.View("index.html")
}

func detail(ctx iris.Context) {
	id := ctx.Params().Get("id")

	result, err := access.Detail(id)
	if err != nil {
		log.Fatal(err)
	}

	ctx.ViewData("body", result)
	ctx.View("article/detail.html")
}

func login(ctx iris.Context) {
	ctx.View("user/login.html")
}

func loginpost(ctx iris.Context) {
	user := models.User{}
	user.UserID = ctx.FormValue("UID")
	user.Password = utils.GetMd5String(ctx.FormValue("PWD"))
	result, err := access.Login(user)
	if err != nil {
		log.Fatal(err)
	}

	if result.ID > 0 {
		id := strconv.Itoa(result.ID)
		// expiration := time.Now()
		// expiration = expiration.AddDate(0, 0, 1)
		ctx.SetCookieKV("id", id, iris.CookieEncode(sc.Encode))
		ctx.SetCookieKV("userid", result.UserID, iris.CookieEncode(sc.Encode))
		ctx.SetCookieKV("username", url.QueryEscape(result.UserName), iris.CookieEncode(sc.Encode))
		ctx.SetCookieKV("token", utils.Encryption(id, result.UserID), iris.CookieEncode(sc.Encode))
	}

	var msg models.Uploador
	msg.Success = result.ID > 0
	ctx.JSON(msg)
}

func logout(ctx iris.Context) {
	ctx.RemoveCookie("id")
	ctx.RemoveCookie("userid")
	ctx.RemoveCookie("username")
	ctx.RemoveCookie("token")

	ctx.Redirect("/")
}

func webpack(ctx iris.Context) {
	ctx.ViewLayout("shared/webpack.html")
	ctx.View("main.html")
}

func addarticle(ctx iris.Context) {
	ctx.ViewLayout("shared/webpack.html")
	ctx.ViewData("id", ctx.Params().Get("id"))
	ctx.ViewData("url", "articleadd")
	ctx.View("article/add.html")
}

func postarticle(ctx iris.Context) {
	var msg models.Uploador
	msg.Success = false

	id, username := utils.GetUser(ctx)
	article := models.Article{
		ID:          ctx.FormValue("Id"),
		NickName:    username,
		UserID:      id,
		Subject:     ctx.FormValue("Subject"),
		Picture:     ctx.FormValue("Picture"),
		PostType:    ctx.FormValue("PostType"),
		Origin:      ctx.FormValue("Origin"),
		Description: ctx.FormValue("Description"),
		Body:        template.HTML(ctx.FormValue("Body"))}

	adding, _ := strconv.ParseBool(ctx.FormValue("Adding"))
	if adding {
		result, err := access.Post(article)
		if err != nil {
			log.Fatal(err)
		}
		msg.Success = result
	} else {
		result, err := access.Update(article)
		if err != nil {
			log.Fatal(err)
		}
		msg.Success = result
	}

	ctx.JSON(msg)
}

func articlelist(ctx iris.Context) {
	id, _ := utils.GetUser(ctx)
	size, _ := strconv.Atoi(ctx.Params().Get("size"))
	page, _ := strconv.Atoi(ctx.Params().Get("page"))
	result, err := access.UserArticle(id, page, size)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(result)
}

func getarticle(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result, err := access.GetArticle(id)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(result)
}

func delarticle(ctx iris.Context) {
	id := ctx.Params().Get("id")
	uid, _ := utils.GetUser(ctx)
	result, err := access.DelArticle(uid, id)
	if err != nil {
		log.Fatal(err)
	}

	var msg models.Uploador
	msg.Success = result
	ctx.JSON(msg)
}

func upload(ctx iris.Context) {
	t := time.Now()
	dir := t.Format("20060102")
	var result models.Uploador
	file, info, err := ctx.FormFile("image")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()
	ext := utils.ExtensionName(info.Filename)
	path := "/upload/" + dir + "/"
	err = os.MkdirAll("./public"+path, os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

	path += utils.UniqueID() + strings.ToLower(ext)
	out, err := os.OpenFile("./public/"+path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer out.Close()
	io.Copy(out, file)

	result.Success = true
	result.Path = path
	ctx.JSON(result)
}

func reg(ctx iris.Context) {
	ctx.View("user/reg.html")
}

func regpost(ctx iris.Context) {
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

func accountlist(ctx iris.Context) {
	size, _ := strconv.Atoi(ctx.Params().Get("size"))
	page, _ := strconv.Atoi(ctx.Params().Get("page"))
	result, err := access.AccountList(page, size)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(result)
}
