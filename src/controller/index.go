package controller

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"github.com/iissy/goweb/src/cli"
	"github.com/iissy/goweb/src/model"
	"github.com/iissy/goweb/src/mosaic"
	"github.com/iissy/goweb/src/utils"
	"github.com/kataras/iris/v12"
	"image"
	"image/jpeg"
	"math"
	"sort"
	"strconv"
	"time"
)

func Index(ctx iris.Context) {
	l := make(chan []*model.Link)
	a := make(chan []*model.Article)
	c := make(chan []*model.CusLink)

	go func() {
		rsp := new(model.LinkList)
		err := cli.Call("IndexLinks", true, rsp)
		utils.WriteErrorLog(ctx.Path(), err)
		l <- rsp.List
	}()

	go func() {
		rsp := new(model.ArticleList)
		err := cli.Call("TopArticles", true, rsp)
		utils.WriteErrorLog(ctx.Path(), err)
		a <- rsp.List
	}()

	go func() {
		rsp := new(model.CusLinkList)
		err := cli.Call("TopCusLinks", true, rsp)
		utils.WriteErrorLog(ctx.Path(), err)
		c <- rsp.List
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
	rspl := new(model.LinkList)
	err := cli.Call("ListLinks", id, rspl)
	utils.WriteErrorLog(ctx.Path(), err)
	links := rspl.List

	rspc := new(model.CusLinkList)
	err = cli.Call("ListCusLinksByCatId", id, rspc)
	utils.WriteErrorLog(ctx.Path(), err)
	cuslinks := rspc.List

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
	rsp := new(model.CusLinkList)
	err := cli.Call("ListCusLinks", true, rsp)
	utils.WriteErrorLog(ctx.Path(), err)
	result := rsp.List

	ctx.ViewData("title", "网络文摘")
	ctx.ViewData("body", result)
	ctx.View("cuslinks.html")
}

func ListArticles(ctx iris.Context) {
	rsp := new(model.ArticleList)
	err := cli.Call("ListArticles", true, rsp)
	utils.WriteErrorLog(ctx.Path(), err)
	result := rsp.List

	ctx.ViewData("title", "文章列表")
	ctx.ViewData("body", result)
	ctx.View("articles.html")
}

func Detail(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result := new(model.Article)
	err := cli.Call("GetArticle", id, result)
	if ok := utils.WriteErrorLog(ctx.Path(), err); ok {
		ctx.NotFound()
		return
	}

	if result == nil {
		ctx.NotFound()
		return
	}

	go cli.Call("UpdateArticleVisited", id, true)
	ctx.ViewData("title", result.Title)
	ctx.ViewData("body", result)
	ctx.View("detail.html")
}

func Payme(ctx iris.Context) {
	result := new(int64)
	err := cli.Call("LinkVisitedCount", true, result)
	utils.WriteErrorLog(ctx.Path(), err)

	ctx.ViewData("title", "打赏站长")
	ctx.ViewData("body", result)
	ctx.View("payme.html")
}

func GetLinkUrl(ctx iris.Context) {
	id := ctx.Params().Get("id")
	url := new(string)
	err := cli.Call("GetLinkUrl", id, url)
	utils.WriteErrorLog(ctx.Path(), err)

	if len(*url) > 0 {
		go cli.Call("UpdateLinkVisited", id, true)
		ctx.Redirect(*url)
	} else {
		ctx.NotFound()
	}
}

func GetCusLinkUrl(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	utils.WriteErrorLog(ctx.Path(), err)

	url := new(string)
	err = cli.Call("GetCusLinkUrl", id, url)
	utils.WriteErrorLog(ctx.Path(), err)

	if len(*url) > 0 {
		go cli.Call("UpdateCusLinkVisited", id, true)
		ctx.Redirect(*url)
	} else {
		ctx.NotFound()
	}
}

func Choose(ctx iris.Context) {
	ctx.ViewData("title", "选择图片")
	ctx.View("choose.html")
}

func Mosaic(ctx iris.Context) {
	t0 := time.Now()
	file, _, _ := ctx.FormFile("image")
	defer file.Close()

	original, _, _ := image.Decode(file)
	bounds := original.Bounds()
	db := mosaic.CloneTilesDB()
	tileSize := int(math.Ceil(math.Sqrt(float64(bounds.Max.X) * float64(bounds.Max.Y) / 3333.0)))
	// fan-out
	c1 := mosaic.Cut(original, &db, tileSize, bounds.Min.X, bounds.Min.Y, bounds.Max.X/2, bounds.Max.Y/2)
	c2 := mosaic.Cut(original, &db, tileSize, bounds.Max.X/2, bounds.Min.Y, bounds.Max.X, bounds.Max.Y/2)
	c3 := mosaic.Cut(original, &db, tileSize, bounds.Min.X, bounds.Max.Y/2, bounds.Max.X/2, bounds.Max.Y)
	c4 := mosaic.Cut(original, &db, tileSize, bounds.Max.X/2, bounds.Max.Y/2, bounds.Max.X, bounds.Max.Y)
	// fan-in
	c := mosaic.Combine(bounds, c1, c2, c3, c4)
	buf1 := new(bytes.Buffer)
	jpeg.Encode(buf1, original, nil)
	originalStr := base64.StdEncoding.EncodeToString(buf1.Bytes())

	t1 := time.Now()
	images := map[string]string{
		"original": originalStr,
		"mosaic":   <-c,
		"duration": fmt.Sprintf("%v ", t1.Sub(t0)),
	}

	ctx.ViewData("title", "图片打马赛克")
	ctx.ViewData("body", images)
	ctx.View("mosaic.html")
}
