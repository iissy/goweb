package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iissy/goweb/src/cli"
	"github.com/iissy/goweb/src/model"
	"github.com/iissy/goweb/src/utils"
	"strconv"
	"time"
)

func GetCusLink(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if ok := utils.WriteErrorLog(ctx.FullPath(), err); ok {
		fmt.Print(0)
	}

	result := new(model.CusLink)
	err = cli.Call("GetCusLink", id, result)
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

	result := new(model.CusLinkList)
	req := model.SearchPager{Pager: model.Pager{Page: page, Size: size}, Search: search}
	err = cli.Call("GetCusLinkList", req, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func DeleteCusLink(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	utils.WriteErrorLog(ctx.FullPath(), err)

	result := new(int64)
	err = cli.Call("DeleteCusLink", id, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func SaveCusLink(ctx *gin.Context) {
	cuslink := new(model.CusLink)
	err := ctx.BindJSON(&cuslink)
	if ok := utils.WriteErrorLog(ctx.FullPath(), err); ok {
		fmt.Print(0)
	}

	cat := new(model.LinkCat)
	err = cli.Call("GetLinkCat", cuslink.Catid, cat)
	utils.WriteErrorLog(ctx.FullPath(), err)

	cuslink.LinkType = cat.CatName
	cuslink.Adddate = time.Now().Format("2006-01-02 15:04:05")
	cuslink.Updatedate = time.Now().Format("2006-01-02 15:04:05")

	result := new(int64)
	err = cli.Call("SaveCusLink", cuslink, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}
