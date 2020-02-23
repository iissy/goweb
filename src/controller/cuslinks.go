package controller

import (
	"github.com/kataras/iris"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
	"strconv"
	"time"
)

func GetCusLink(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	if ok := utils.WriteErrorLog(ctx, err); ok {
		ctx.JSON(0)
	}

	result, err := domain.GetCusLink(id)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}

func GetCusLinkList(ctx iris.Context) {
	size, err := strconv.Atoi(ctx.Params().Get("size"))
	if err != nil {
		size = 10
	}

	page, err := strconv.Atoi(ctx.Params().Get("page"))
	if err != nil {
		page = 1
	}

	search := new(model.Search)
	err = ctx.ReadJSON(&search)
	utils.WriteErrorLog(ctx, err)

	result, err := domain.GetCusLinkList(page, size, search)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}

func DeleteCusLink(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	utils.WriteErrorLog(ctx, err)

	result, err := domain.DeleteCusLink(id)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}

func SaveCusLink(ctx iris.Context) {
	cuslink := new(model.CusLink)
	err := ctx.ReadJSON(&cuslink)
	if ok := utils.WriteErrorLog(ctx, err); ok {
		ctx.JSON(0)
	}

	cat, err := domain.GetLinkCat(cuslink.Catid)
	utils.WriteErrorLog(ctx, err)

	cuslink.LinkType = cat.CatName
	cuslink.Adddate = time.Now().Format("2006-01-02 15:04:05")
	cuslink.Updatedate = time.Now().Format("2006-01-02 15:04:05")
	result, err := domain.SaveCusLink(cuslink)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}
