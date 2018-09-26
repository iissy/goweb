package utils

import (
	"net/http"
	"net/url"
	"strconv"
)

// GetUser is yes
func GetUser(r *http.Request) (int, string) {
	if Check(r) {
		idCookie, _ := r.Cookie("id")
		unmCookie, _ := r.Cookie("username")
		id, _ := strconv.Atoi(idCookie.Value)
		nickname, _ := url.QueryUnescape(unmCookie.Value)
		return id, nickname
	}
	return -1, ""
}

// Check is yes
func Check(r *http.Request) bool {
	idCookie, _ := r.Cookie("id")
	uidCookie, _ := r.Cookie("userid")
	unmCookie, _ := r.Cookie("username")
	tokenCookie, _ := r.Cookie("token")
	if idCookie == nil || uidCookie == nil || unmCookie == nil && tokenCookie == nil {
		return false
	}

	token := Encryption(idCookie.Value, uidCookie.Value)
	if tokenCookie.Value == token {
		return true
	}
	return false
}
