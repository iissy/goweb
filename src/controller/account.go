package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/utils"
	"strconv"
)

func GetAccount(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if ok := utils.WriteErrorLog(ctx.FullPath(), err); ok {
		fmt.Print(0)
	}

	result, err := domain.GetAccount(id)
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

	result, err := domain.GetrAccountList(page, size)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}
