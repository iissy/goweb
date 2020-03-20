package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kataras/golog"
	"os"
	"path"
	"strings"
	"time"
)

func Upload(ctx *gin.Context) {
	result := struct {
		Message  string `json:"msg"`
		Uploaded int    `json:"uploaded"`
		Url      string `json:"url"`
	}{"", 0, ""}

	file, err := ctx.FormFile("upload")
	if err != nil {
		golog.Error(err)
		ctx.JSON(500, result)
	}

	//file, info, err := ctx.FormFile("upload")
	//if err != nil {
	//	golog.Error(err)
	//	ctx.JSON(result)
	//}

	dir := "/upload/" + time.Now().Format("20060102") + "/"
	err = os.MkdirAll("./public"+dir, os.ModePerm)
	if err != nil {
		golog.Error(err)
		ctx.JSON(500, result)
	}

	url := dir + uuid.New().String() + strings.ToLower(path.Ext(file.Filename))
	//out, err := os.OpenFile("./public/"+url, os.O_WRONLY|os.O_CREATE, 0666)
	//if err != nil {
	//	golog.Error(err)
	//	ctx.JSON(result)
	//}

	// Upload the file to specific dst.
	dst := "./public/" + url
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		golog.Error(err)
		ctx.JSON(500, result)
	}

	//defer file.Close()
	//defer out.Close()
	//io.Copy(out, file)

	result.Uploaded = 1
	result.Url = url
	ctx.JSON(200, result)
}
