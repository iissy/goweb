package utils

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gorilla/securecookie"
	"math/rand"
	"time"
)

const code = "0123456789ABCDEFGHIJKLMNOPQRSTUVXWYZabcdefghijklmnopqrstuvxwyz-*"

var (
	hashKey   = []byte("the-big-and-secret-fash-key-here")
	blockKey  = []byte("lot-secret-of-characters-big-too")
	SC        = securecookie.New(hashKey, blockKey)
	ASYUSERID = "asy-user-id"
	ASYTOKEN  = "asy-token"
)

func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func randomString(size int, max int) string {
	rand.Seed(time.Now().UnixNano())
	buffer := make([]byte, size, size)
	for i := 0; i < size; i++ {
		buffer[i] = code[rand.Intn(max)]
	}
	return string(buffer[:size])
}

func Random62String(size int) string {
	return randomString(size, 62)
}
