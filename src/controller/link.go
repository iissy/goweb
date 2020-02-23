package controller

import (
	"github.com/google/uuid"
	"github.com/kataras/iris"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
	"strconv"
	"time"
)

func GetLink(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result, err := domain.GetLink(id)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}

func GetLinkList(ctx iris.Context) {
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

	result, err := domain.GetLinkList(page, size, search)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}

func DeleteLink(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result, err := domain.DeleteLink(id)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}

func SaveLink(ctx iris.Context) {
	link := new(model.Link)
	err := ctx.ReadJSON(&link)
	if ok := utils.WriteErrorLog(ctx, err); ok {
		ctx.JSON(0)
	}

	cat, err := domain.GetLinkCat(link.Catid)
	utils.WriteErrorLog(ctx, err)

	if len(link.Id) < 10 {
		link.Id = uuid.New().String()
	}
	link.LinkType = cat.CatName
	link.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	result, err := domain.SaveLink(link)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}
