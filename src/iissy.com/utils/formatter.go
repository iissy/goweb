package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"html/template"
	"io"
	"math"
	"strconv"
	"time"

	"iissy.com/models"
)

// ExtensionName 截取字符串 start 起点下标
func ExtensionName(str string) string {
	rs := []rune(str)
	var start int
	for i := range rs {
		if string(rs[i]) == "." {
			start = i
		}
	}

	return string(rs[start:])
}

// Substr 截取字符串 start 起点下标
func Substr(str string, start int) string {
	rs := []rune(str)
	return string(rs[start:])
}

// Substr2 截取字符串 start 起点下标 end 终点下标(不包括)
func Substr2(str string, start int, length int) string {
	rs := []rune(str)
	rl := len(rs)
	end := 0

	if start < 0 {
		start = rl - 1 + start
	}
	end = start + length

	if start > end {
		start, end = end, start
	}

	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}

	return string(rs[start:end])
}

// GetMd5String 生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// UniqueID 生成Guid字串
func UniqueID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5String(base64.URLEncoding.EncodeToString(b))
}

// Encryption 生成Guid字串
func Encryption(params ...string) string {
	token := "iamjimmyiuokyesfuck"
	for _, param := range params {
		token = GetMd5String(param + token)
	}
	return token
}

// DaysString 时间格式化
func DaysString(str string) string {
	layout := "2006-01-02 15:04:05"
	t, _ := time.Parse(layout, str)
	day := t.Format("2006-01-02 15:04:05")
	return day
}

// PageChanging 翻页
func PageChanging(pageArgs models.PageArgs) template.HTML {
	var pages string
	size := float64(pageArgs.PageSize)
	page := pageArgs.PageNumber
	pageCount := int(math.Ceil(float64(pageArgs.TotalCount) / size))
	if page <= 1 || pageCount <= 1 {
		pages += "<li class='prev disabled'><a href='#' title='Prev'>&laquo;</a></li>"
	} else {
		pages += "<li class='prev'><a href='/article/list/1' title='Prev'>&laquo;</a></li>"
	}
	start := 1
	end := pageCount
	if page > 3 {
		start = page - 3
	}
	if (pageCount - page) > 3 {
		end = page + 3
	}

	for i := start; i <= end; i++ {
		if i == page {
			pages += "<li class='active'><a href='/article/list/" + strconv.Itoa(i) + "'>" + strconv.Itoa(i) + "</a></li>"
		} else {
			pages += "<li><a href='/article/list/" + strconv.Itoa(i) + "'>" + strconv.Itoa(i) + "</a></li>"
		}
	}

	if page < pageCount && pageCount > 1 {
		pages += "<li class='next'><a href='/article/list/" + strconv.Itoa(pageCount) + "' title='Next'>&raquo;</a></li>"
	} else {
		pages += "<li class='next disabled'><a href='#' title='Next'>&raquo;</a></li>"
	}
	return template.HTML(pages)
}
