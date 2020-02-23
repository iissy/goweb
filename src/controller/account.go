package controller

import (
	"github.com/kataras/iris"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/utils"
	"strconv"
)

func GetAccount(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	if ok := utils.WriteErrorLog(ctx, err); ok {
		ctx.JSON(0)
	}

	result, err := domain.GetAccount(id)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}

func GetAccountList(ctx iris.Context) {
	size, err := strconv.Atoi(ctx.Params().Get("size"))
	if err != nil {
		size = 10
	}

	page, err := strconv.Atoi(ctx.Params().Get("page"))
	if err != nil {
		page = 1
	}

	result, err := domain.GetrAccountList(page, size)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}
