package utils

import (
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"io"
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
