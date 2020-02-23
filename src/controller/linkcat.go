package controller

import (
	"github.com/kataras/iris"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/utils"
)

func GetCatOptions(ctx iris.Context) {
	result, err := domain.GetCatOptions()
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}
