package controller

import (
	"github.com/gin-gonic/gin"
	"hrefs.cn/src/cli"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
)

func GetCatOptions(ctx *gin.Context) {
	result := new(model.LinkCatList)
	err := cli.Call("GetCatOptions", true, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result.List)
}
