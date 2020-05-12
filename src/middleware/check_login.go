package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/iissy/goweb/src/cli"
	"github.com/iissy/goweb/src/utils"
	"github.com/juju/errors"
	"log"
)

func CheckLogin(ctx *gin.Context) {
	id := ParseHeadOrCookie(ctx, utils.ASYUSERID)
	token := ParseHeadOrCookie(ctx, utils.ASYTOKEN)

	if len(id) <= 0 || len(token) <= 0 {
		log.Printf("miss id or token")
		ctx.Abort()
		return
	}

	v := new(string)
	err := cli.Call("GetToken", id, v)
	if err != nil {
		log.Printf(errors.ErrorStack(err))
		ctx.Abort()
		return
	}

	if *v != token {
		ctx.Abort()
	}
}

func ParseHeadOrCookie(ctx *gin.Context, k string) string {
	v := ctx.GetHeader(k)
	if len(v) == 0 {
		v, _ = ctx.Cookie(k)
	}
	return v
}
