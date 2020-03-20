package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
	"strconv"
	"time"
)

func GetLink(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := domain.GetLink(id)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func GetLinkList(ctx *gin.Context) {
	size, err := strconv.Atoi(ctx.Param("size"))
	if err != nil {
		size = 10
	}

	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil {
		page = 1
	}

	search := new(model.Search)
	err = ctx.BindJSON(&search)
	utils.WriteErrorLog(ctx.FullPath(), err)

	result, err := domain.GetLinkList(page, size, search)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func DeleteLink(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := domain.DeleteLink(id)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func SaveLink(ctx *gin.Context) {
	link := new(model.Link)
	err := ctx.BindJSON(&link)
	if ok := utils.WriteErrorLog(ctx.FullPath(), err); ok {
		fmt.Print(0)
	}

	cat, err := domain.GetLinkCat(link.Catid)
	utils.WriteErrorLog(ctx.FullPath(), err)

	if len(link.Id) < 10 {
		link.Id = uuid.New().String()
	}
	link.LinkType = cat.CatName
	link.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	result, err := domain.SaveLink(link)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}
