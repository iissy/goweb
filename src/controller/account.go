package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iissy/goweb/src/cli"
	"github.com/iissy/goweb/src/model"
	"github.com/iissy/goweb/src/utils"
	"strconv"
)

func GetAccount(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if ok := utils.WriteErrorLog(ctx.FullPath(), err); ok {
		fmt.Print(0)
	}

	result := new(model.Account)
	err = cli.Call("GetAccount", id, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func GetAccountList(ctx *gin.Context) {
	size, err := strconv.Atoi(ctx.Param("size"))
	if err != nil {
		size = 10
	}

	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil {
		page = 1
	}

	result := new(model.AccountList)
	req := model.Pager{Page: page, Size: size}
	err = cli.Call("GetAccountList", req, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}
