package controller

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gorilla/securecookie"
	"github.com/kataras/iris"
	"iissy.com/models"
	"iissy.com/utils"
)

var (
	hashKey  = []byte("the-big-and-secret-fash-key-here")
	blockKey = []byte("lot-secret-of-characters-big-too")
	sc       = securecookie.New(hashKey, blockKey)
)

// Before 前置控制器
func Before(ctx iris.Context) {
	id, _, username := utils.GetUser(ctx)
	ctx.ViewData("islogin", id > 0)
	ctx.ViewData("username", username)
	ctx.Next()
}

// Upload 上传
func Upload(ctx iris.Context) {
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
