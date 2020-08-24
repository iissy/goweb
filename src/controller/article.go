package controller

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iissy/goweb/src/cli"
	"github.com/iissy/goweb/src/model"
	"github.com/iissy/goweb/src/utils"
)

func GetArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	result := new(model.Article)
	err := cli.Call("GetArticle", id, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func GetArticleList(ctx *gin.Context) {
	size, err := strconv.Atoi(ctx.Param("size"))
	if err != nil {
		size = 10
	}

	page, err := strconv.Atoi(ctx.Param("page"))
	if err != nil {
		page = 1
	}

	search := new(model.Search)
	err = ctx.BindJSON(&search)
	utils.WriteErrorLog(ctx.FullPath(), err)

	req := model.SearchPager{Pager: model.Pager{Page: page, Size: size}, Search: search}
	result := new(model.ArticleList)
	err = cli.Call("GetArticleList", req, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")

	result := new(int64)
	err := cli.Call("DeleteArticle", id, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func SaveArticle(ctx *gin.Context) {
	article := new(model.Article)
	err := ctx.BindJSON(&article)
	if ok := utils.WriteErrorLog(ctx.FullPath(), err); ok {
		fmt.Print(0)
	}

	//将HTML标签全转换成小写
	brief := string(article.Body)
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	brief = re.ReplaceAllStringFunc(brief, strings.ToLower)
	//去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	brief = re.ReplaceAllString(brief, "")
	//去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	brief = re.ReplaceAllString(brief, "")
	//去除所有尖括号内的HTML代码，并换成换行符
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	brief = re.ReplaceAllString(brief, "\n")
	//去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	brief = re.ReplaceAllString(brief, "\n")
	article.Brief = brief

	article.CreateTime = time.Now()

	result := new(int64)
	err = cli.Call("SaveArticle", article, result)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}
