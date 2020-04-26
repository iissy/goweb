package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/iissy/goweb/src/cli"
	"github.com/iissy/goweb/src/model"
	"github.com/iissy/goweb/src/utils"
)

func GetCatOptions(ctx *gin.Context) {
	result := new(model.LinkCatList)
	err := cli.Call("GetCatOptions", true, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result.List)
}
