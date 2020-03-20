package controller

import (
	"github.com/gin-gonic/gin"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/utils"
)

func GetCatOptions(ctx *gin.Context) {
	result, err := domain.GetCatOptions()
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}
