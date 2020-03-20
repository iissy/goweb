package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
	"strconv"
	"time"
)

func GetCusLink(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if ok := utils.WriteErrorLog(ctx.FullPath(), err); ok {
		fmt.Print(0)
	}

	result, err := domain.GetCusLink(id)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func GetCusLinkList(ctx *gin.Context) {
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

	result, err := domain.GetCusLinkList(page, size, search)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func DeleteCusLink(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	utils.WriteErrorLog(ctx.FullPath(), err)

	result, err := domain.DeleteCusLink(id)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func SaveCusLink(ctx *gin.Context) {
	cuslink := new(model.CusLink)
	err := ctx.BindJSON(&cuslink)
	if ok := utils.WriteErrorLog(ctx.FullPath(), err); ok {
		fmt.Print(0)
	}

	cat, err := domain.GetLinkCat(cuslink.Catid)
	utils.WriteErrorLog(ctx.FullPath(), err)

	cuslink.LinkType = cat.CatName
	cuslink.Adddate = time.Now().Format("2006-01-02 15:04:05")
	cuslink.Updatedate = time.Now().Format("2006-01-02 15:04:05")
	result, err := domain.SaveCusLink(cuslink)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}
