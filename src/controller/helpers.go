package controller

import (
	"github.com/google/uuid"
	"github.com/kataras/golog"
	"github.com/kataras/iris"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

func Upload(ctx iris.Context) {
	result := struct {
		Message  string `json:"msg"`
		Uploaded int    `json:"uploaded"`
		Url      string `json:"url"`
	}{"", 0, ""}

	file, info, err := ctx.FormFile("upload")
	if err != nil {
		golog.Error(err)
		ctx.JSON(result)
	}

	dir := "/upload/" + time.Now().Format("20060102") + "/"
	err = os.MkdirAll("./public"+dir, os.ModePerm)
	if err != nil {
		golog.Error(err)
		ctx.JSON(result)
	}

	url := dir + uuid.New().String() + strings.ToLower(path.Ext(info.Filename))
	out, err := os.OpenFile("./public/"+url, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		golog.Error(err)
		ctx.JSON(result)
	}

	defer file.Close()
	defer out.Close()
	io.Copy(out, file)

	result.Uploaded = 1
	result.Url = url
	ctx.JSON(result)
}
