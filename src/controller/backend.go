package controller

import (
	"html/template"
	"log"
	"strconv"

	"github.com/kataras/iris"
	"iissy.com/src/access"
	"iissy.com/src/models"
	"iissy.com/src/utils"
)

// Webpack 授权基础页面
func Webpack(ctx iris.Context) {
	ctx.ViewLayout("shared/webpack.html")
	ctx.View("main.html")
}

// Postarticle 发布文章
func Postarticle(ctx iris.Context) {
	var msg models.Uploador
	msg.Success = false

	id, _, username := utils.GetUser(ctx)
	article := models.Article{
		ID:          ctx.FormValue("Id"),
		NickName:    username,
		UserID:      id,
		Subject:     ctx.FormValue("Subject"),
		Picture:     ctx.FormValue("Picture"),
		PostType:    ctx.FormValue("PostType"),
		Origin:      ctx.FormValue("Origin"),
		Description: ctx.FormValue("Description"),
		Body:        template.HTML(ctx.FormValue("Body"))}

	adding, _ := strconv.ParseBool(ctx.FormValue("Adding"))
	if adding {
		result, err := access.Post(article)
		if err != nil {
			log.Fatal(err)
		}
		msg.Success = result
	} else {
		result, err := access.Update(article)
		if err != nil {
			log.Fatal(err)
		}
		msg.Success = result
	}

	ctx.JSON(msg)
}

// Articlelist 文章列表
func Articlelist(ctx iris.Context) {
	id, _, _ := utils.GetUser(ctx)
	size, _ := strconv.Atoi(ctx.Params().Get("size"))
	page, _ := strconv.Atoi(ctx.Params().Get("page"))
	result, err := access.UserArticle(id, page, size)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(result)
}

// Getarticle 获取文章
func Getarticle(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result, err := access.GetArticle(id)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(result)
}

// Delarticle 删除文章
func Delarticle(ctx iris.Context) {
	id := ctx.Params().Get("id")
	uid, _, _ := utils.GetUser(ctx)
	result, err := access.DelArticle(uid, id)
	if err != nil {
		log.Fatal(err)
	}

	var msg models.Uploador
	msg.Success = result
	ctx.JSON(msg)
}

// Accountlist 用户列表
func Accountlist(ctx iris.Context) {
	size, _ := strconv.Atoi(ctx.Params().Get("size"))
	page, _ := strconv.Atoi(ctx.Params().Get("page"))
	result, err := access.AccountList(page, size)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(result)
}

// Postrole 添加角色
func Postrole(ctx iris.Context) {
	var msg models.Uploador
	msg.Success = false

	id, err := strconv.Atoi(ctx.FormValue("Id"))
	status, err := strconv.Atoi(ctx.FormValue("Status"))
	if err != nil {
		log.Fatal(err)
	}

	role := models.Role{
		ID:       id,
		RoleName: ctx.FormValue("RoleName"),
		Status:   status}

	result, err := access.PostRole(role)
	if err != nil {
		log.Fatal(err)
	}
	msg.Success = result

	ctx.JSON(msg)
}

// Rolelist 角色类别
func Rolelist(ctx iris.Context) {
	id, _, _ := utils.GetUser(ctx)
	size, _ := strconv.Atoi(ctx.Params().Get("size"))
	page, _ := strconv.Atoi(ctx.Params().Get("page"))
	result, err := access.RoleList(id, page, size)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(result)
}

// Getrole 获取角色
func Getrole(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result, err := access.GetRole(id)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(result)
}

// Postfunction 添加权限
func Postfunction(ctx iris.Context) {
	var msg models.Uploador
	msg.Success = false

	id, _ := strconv.Atoi(ctx.FormValue("Id"))
	fun := models.Functionality{
		ID:         id,
		Funname:    ctx.FormValue("Funname"),
		FunType:    ctx.FormValue("FunType"),
		Controller: ctx.FormValue("Controller")}

	result, err := access.PostFunction(fun)
	if err != nil {
		log.Fatal(err)
	}
	msg.Success = result

	ctx.JSON(msg)
}

// Functionlist 权限列表
func Functionlist(ctx iris.Context) {
	id, _, _ := utils.GetUser(ctx)
	size, _ := strconv.Atoi(ctx.Params().Get("size"))
	page, _ := strconv.Atoi(ctx.Params().Get("page"))
	result, err := access.FunctionList(id, page, size)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(result)
}

// Getfunction 获取权限
func Getfunction(ctx iris.Context) {
	id := ctx.Params().Get("id")
	result, err := access.GetFunction(id)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(result)
}

// Functiongroup 角色权限表
func Functiongroup(ctx iris.Context) {
	id, err := strconv.Atoi(ctx.Params().Get("id"))
	if err != nil {
		log.Fatal(err)
	}

	functions, err := access.FunctionGroup()
	if err != nil {
		log.Fatal(err)
	}

	selectedids, err := access.GetRoleFunction(id)
	if err != nil {
		log.Fatal(err)
	}

	ctx.JSON(struct {
		Functions   map[string][]*models.Functionality `json:"functions"`
		Selectedids []int                              `json:"selectedids"`
	}{functions, selectedids})
}

// Mappingpost 权限配置
func Mappingpost(ctx iris.Context) {
	var msg models.Uploador
	msg.Success = false

	fundid, _ := strconv.Atoi(ctx.FormValue("FunId"))
	roleid, _ := strconv.Atoi(ctx.FormValue("RoleId"))
	toggle, _ := strconv.ParseBool(ctx.FormValue("Toggle"))

	mapping := models.RoleFunctionMapping{
		FunID:  fundid,
		RoleID: roleid,
		Toggle: toggle}

	result, err := access.MappingPost(mapping)
	if err != nil {
		log.Fatal(err)
	}
	msg.Success = result

	ctx.JSON(msg)
}
