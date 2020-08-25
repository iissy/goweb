package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
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
		logrus.Error(err)
		ctx.JSON(500, result)
	}

	dir := "/upload/" + time.Now().Format("20060102") + "/"
	err = os.MkdirAll("./public"+dir, os.ModePerm)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(500, result)
	}

	url := dir + uuid.New().String() + strings.ToLower(path.Ext(file.Filename))

	dst := "./public/" + url
	err = ctx.SaveUploadedFile(file, dst)
	if err != nil {
		logrus.Error(err)
		ctx.JSON(500, result)
	}

	result.Uploaded = 1
	result.Url = url
	ctx.JSON(200, result)
}
