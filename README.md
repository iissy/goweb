这是一个由Go语言写的，简单的CMS（内容管理系统），已经更新到 Go 1.13 版本。

## 星星增长趋势
[![Stargazers over time](https://starchart.cc/iissy/goweb.svg)](https://starchart.cc/iissy/goweb)

#### asy 2.0 Released 2020年02月23日
+ 技术大更新

### 技术栈
+ go iris 框架
+ cookie 认证，双重加密
+ 文件目录结构调整
+ 前端路由，后端路径路由与分组路由
+ 中间件实现授权访问
+ 使用 Go Module 方式开发

### 功能点
+ 文章列表显示，详情显示，并支持SEO
+ 前端图片懒加载
+ 登录，注册
+ 后台内容管理，用户管理
+ 图片上传
+ 单页面后台管理
+ 网址管理，网摘管理

## 安装说明
1. 安装mysql
2. 使用hrefs.cn.sql初始化数据库以及数据
3. 修改数据库连接（conf/config.json）

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
+ 账号：jimmy 密码：123456


## 相关网站
+ 程序员网址导航：https://www.hrefs.cn