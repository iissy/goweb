package controller

import (
	"log"
	"net/http"
	"reflect"
	"runtime"
	"strings"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
	"iissy.com/access"
	"iissy.com/cache"
	"iissy.com/models"
	"iissy.com/utils"
)

var memo *cache.Memo

func init() {
	memo = cache.New(roleGetFunction)
	return
}

// 根据角色获取所有权限
func roleGetFunction(roleid int) ([]string, error) {
	funclist, err := access.GetFunctionNames(roleid)
	return funclist, err
}

// BasicAuth 是登录认证，用户分权限管理
func BasicAuth(h context.Handler) context.Handler {
	return func(ctx iris.Context) {
		if ok := utils.Check(ctx); ok {
			name := runtime.FuncForPC(reflect.ValueOf(h).Pointer()).Name()
			index := strings.LastIndex(name, ".")
			name = strings.ToLower(name[index+1:])

			log.Print(name)
			_, roleid, _ := utils.GetUser(ctx)
			funclist, err := memo.Get(roleid)
			if err != nil {
				log.Print(err)
			}

			flag := false
			for _, item := range funclist {
				if item == name {
					flag = true
					break
				}
			}

			if flag {
				h(ctx)
			} else {
				ctx.JSON(models.Author{Success: false, Message: "没有权限"})
			}
		} else {
			http.Error(ctx.ResponseWriter(), http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		}
	}
}
