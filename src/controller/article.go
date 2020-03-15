package controller

import (
	"github.com/kataras/iris"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetArticle(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result, err := domain.GetArticle(id)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}

func GetArticleList(ctx iris.Context) {
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

	result, err := domain.GetArticleList(page, size, search)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}

func DeleteArticle(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result, err := domain.DeleteArticle(id)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}

func SaveArticle(ctx iris.Context) {
	article := new(model.Article)
	err := ctx.ReadJSON(&article)
	if ok := utils.WriteErrorLog(ctx, err); ok {
		ctx.JSON(0)
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

	article.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	result, err := domain.SaveArticle(article)
	utils.WriteErrorLog(ctx, err)

	ctx.JSON(result)
}
