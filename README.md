这是一个由Go语言写的，简单的CMS（内容管理系统），已经更新到 Go 1.13 版本。

## 星星增长趋势
[![Stargazers over time](https://starchart.cc/iissy/goweb.svg)](https://starchart.cc/iissy/goweb)

#### Asy 1.8 Released 2019年11月19日
+ 解决图片上传问题

### 技术栈
+ 改用 iris 框架
+ cookie 认证，双重加密
+ 文件目录结构调整
+ 引入 webpack + vue-router + axios
+ 添加角色权限管理
+ 前端路由，后端路径路由与分组路由
+ 注入缓存模块
+ 中间件实现授权访问
+ 使用 Go Module 方式开发

### 功能点
+ 包括内容列表显示，详情显示，并支持SEO
+ 前端图片懒加载
+ 登录，注册
+ 后台内容管理，用户管理，角色权限管理
+ 图片上传
+ 单页面后台管理

## 安装说明
1. 安装mysql
2. 使用iissy.com.sql初始化数据库以及数据
3. 修改数据库连接（src/iissy.com/utils/config.go）

## 使用代码
由于大陆网络无法下载google的包，使用七牛公司的代理，下载包前执行命令
1. go env -w GOPROXY=https://goproxy.cn,direct

## 组件下载
1. go get github.com/go-sql-driver/mysql
2. go get github.com/kataras/iris
3. go get github.com/gorilla/securecookie
4. 运行命令启动# go run main.go

## 使用说明
+ 请确保80端口没有被别的服务（比如iis）占用，然后在浏览器中输入：http://localhost
+ 管理员账号：jimmy 密码：123456
+ 普通账号：pinbor 密码：123456

## 首页
![markdown](https://github.com/iissy/goweb/blob/master/public/home.png "首页图片")

## 文章
![markdown](https://github.com/iissy/goweb/blob/master/public/art.png "文章图片")

## 发布
![markdown](https://github.com/iissy/goweb/blob/master/public/add.png "发布图片")

## 列表
![markdown](https://github.com/iissy/goweb/blob/master/public/list.png "列表图片")

## 角色
![markdown](https://github.com/iissy/goweb/blob/master/public/role.png "角色图片")

## 权限
![markdown](https://github.com/iissy/goweb/blob/master/public/fun.png "权限图片")

## 权限配置
![markdown](https://github.com/iissy/goweb/blob/master/public/map.png "权限配置图片")


## 相关网站
+ 程序员网址导航：https://www.hrefs.cn
+ 技术文档：https://www.hrefs.cn/article/Go-iris-webpack-vue-router-axios-CMS
