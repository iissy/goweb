package controller

import (
	"github.com/kataras/iris"
	"hrefs.cn/src/cli"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
	"sort"
	"strconv"
)

func Index(ctx iris.Context) {
	l := make(chan []*model.Link)
	a := make(chan []*model.Article)
	c := make(chan []*model.CusLink)

	go func() {
		rsp := new(model.LinkItems)
		err := cli.Call("IndexLinks", 0, rsp)
		utils.WriteErrorLog(ctx.Path(), err)
		l <- rsp.Items
	}()

	go func() {
		rsp := new(model.ArticleItems)
		err := cli.Call("TopArticles", 0, rsp)
		utils.WriteErrorLog(ctx.Path(), err)
		a <- rsp.Items
	}()

	go func() {
		rsp := new(model.CusLinkItems)
		err := cli.Call("TopCusLinks", 0, rsp)
		utils.WriteErrorLog(ctx.Path(), err)
		c <- rsp.Items
	}()

	result := new(model.Index)
	linkMap := make(map[string][]*model.Link)
	linkCountMap := make(map[string]int)
	catMap := make(map[string]string)

	links := <-l
	for _, link := range links {
		linkMap[link.LinkType] = append(linkMap[link.LinkType], link)
		linkCountMap[link.LinkType] = len(linkMap[link.LinkType])
		catMap[link.LinkType] = link.Catid
	}

	groups := make(model.OneGroups, 0)
	for key, val := range catMap {
		item := new(model.OneGroup)
		item.CatId = val
		item.Name = key
		item.Size = linkCountMap[key]
		item.Items = linkMap[key]
		groups = append(groups, item)
	}
	sort.Sort(groups)

	result.Common = linkMap["公共"]
	result.Info = linkMap["资讯"]
	result.Articles = <-a
	result.CusLinks = <-c

	result.Other = setter([]string{"实用工具", "软件", "教程"}, groups)
	result.Backend = setter([]string{"NodeJS", "PHP", "DotNet", "Golang", "Java", "Python", "其他"}, groups)
	result.Frontend = setter([]string{"CSS", "JQuery", "Charts", "Vue", "前端框架", "富文本编辑器", "打包构建"}, groups)
	result.Hot = setter([]string{"架构师", "人工智能", "区块链", "大数据", "数据库", "运维工具", "协同工具"}, groups)

	ctx.ViewLayout(iris.NoLayout)
	ctx.ViewData("body", result)
	ctx.View("index.html")
}

func setter(list []string, groups model.OneGroups) *model.LinkGroup {
	result := new(model.LinkGroup)
	newGroups := make(model.OneGroups, 0)
	for _, g := range groups {
		for _, l := range list {
			if l == g.Name {
				newGroups = append(newGroups, g)
				break
			}
		}
	}

	if len(newGroups) > 0 {
		result.Active = newGroups[0]
		result.Rests = newGroups[1:]
	}

	return result
}

func ListLinks(ctx iris.Context) {
	id := ctx.Params().Get("id")
	links, err := domain.ListLinks(id)
	utils.WriteErrorLog(ctx.Path(), err)

	cuslinks, err := domain.ListCusLinksByCatId(id)
	utils.WriteErrorLog(ctx.Path(), err)

	if links == nil || len(links) == 0 {
		ctx.NotFound()
	} else {
		ctx.ViewData("title", links[0].LinkType+" - 网址大全")
		ctx.ViewData("links", links)
		ctx.ViewData("cuslinks", cuslinks)
		ctx.View("links.html")
	}
}

func ListCusLinks(ctx iris.Context) {
	result, err := domain.ListCusLinks()
	utils.WriteErrorLog(ctx.Path(), err)

	ctx.ViewData("title", "网络文摘")
	ctx.ViewData("body", result)
	ctx.View("cuslinks.html")
}

func ListArticles(ctx iris.Context) {
	result, err := domain.ListArticles()
	utils.WriteErrorLog(ctx.Path(), err)

	ctx.ViewData("title", "文章列表")
	ctx.ViewData("body", result)
	ctx.View("articles.html")
}

func Detail(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result, err := domain.GetArticle(id)
	if ok := utils.WriteErrorLog(ctx.Path(), err); ok {
		ctx.NotFound()
		return
	}

	if result == nil {
		ctx.NotFound()
		return
	}

	go domain.UpdateArticleVisited(id)
	ctx.ViewData("title", result.Title)
	ctx.ViewData("body", result)
	ctx.View("detail.html")
}

func Payme(ctx iris.Context) {
	result, err := domain.LinkVisitedCount()
	utils.WriteErrorLog(ctx.Path(), err)

	ctx.ViewData("title", "打赏站长")
	ctx.ViewData("body", result)
	ctx.View("payme.html")
}

func GetLinkUrl(ctx iris.Context) {
	id := ctx.Params().Get("id")
	url, err := domain.GetLinkUrl(id)
	utils.WriteErrorLog(ctx.Path(), err)

	if len(url) > 0 {
		go domain.UpdateLinkVisited(id)
		ctx.Redirect(url)
	} else {
		ctx.NotFound()
	}
}

func GetCusLinkUrl(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	utils.WriteErrorLog(ctx.Path(), err)

	url, err := domain.GetCusLinkUrl(id)
	utils.WriteErrorLog(ctx.Path(), err)

	if len(url) > 0 {
		go domain.UpdateCusLinkVisited(id)
		ctx.Redirect(url)
	} else {
		ctx.NotFound()
	}
}
