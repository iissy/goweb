package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/model"
	"hrefs.cn/src/utils"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func GetArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := domain.GetArticle(id)
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

	result, err := domain.GetArticleList(page, size, search)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}

func DeleteArticle(ctx *gin.Context) {
	id := ctx.Param("id")
	result, err := domain.DeleteArticle(id)
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

	article.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	result, err := domain.SaveArticle(article)
	utils.WriteErrorLog(ctx.FullPath(), err)

	ctx.JSON(200, result)
}
