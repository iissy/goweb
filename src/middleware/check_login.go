package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/juju/errors"
	"hrefs.cn/src/redis"
	"hrefs.cn/src/utils"
	"log"
)

func CheckLogin(ctx *gin.Context) {
	id := ParseHeadOrCookie(ctx, utils.ASYUSERID)
	token := ParseHeadOrCookie(ctx, utils.ASYTOKEN)

	if len(id) <= 0 || len(token) <= 0 {
		log.Printf("miss id or token")
		return
	}

	v, err := redis.Get(id)
	if err != nil {
		log.Printf(errors.ErrorStack(err))
		return
	}

	if v == token {
		ctx.Next()
	}
}

func ParseHeadOrCookie(ctx *gin.Context, k string) string {
	v := ctx.GetHeader(k)
	if len(v) == 0 {
		v, _ = ctx.Cookie(k)
	}
	return v
}
