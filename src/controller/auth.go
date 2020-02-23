package controller

import (
	"github.com/kataras/iris"
	"hrefs.cn/src/domain"
	"hrefs.cn/src/model"
	"hrefs.cn/src/redis"
	"hrefs.cn/src/utils"
	"time"
)

func Login(ctx iris.Context) {
	user := new(model.Account)
	err := ctx.ReadJSON(&user)
	if ok := utils.WriteErrorLog(ctx, err); ok {
		ctx.JSON(0)
	}

	user.Password = utils.GetMd5String(user.Password)
	user.LastLoginDate = time.Now().Format("2006-01-02 15:04:05")
	result, err := domain.Login(user)
	if ok := utils.WriteErrorLog(ctx, err); ok {
		ctx.JSON(0)
	}

	if result.ID > 0 {
		token := utils.Random62String(64)
		ctx.Header(utils.ASYUSERID, result.UserId)
		ctx.Header(utils.ASYTOKEN, token)
		ctx.SetCookieKV(utils.ASYUSERID, result.UserId, iris.CookieEncode(utils.SC.Encode))
		ctx.SetCookieKV(utils.ASYTOKEN, token, iris.CookieEncode(utils.SC.Encode))
		err = redis.Set(result.UserId, token)
		utils.WriteErrorLog(ctx, err)
	}

	ctx.JSON(result)
}
