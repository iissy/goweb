这是一个由Go语言写的，简单的内容管理系统CMS，2017年06月30日启动项目，至今仍在更新维护，见下面版本发布时间。

### 功能点
+ 包括前端显示，SEO支持
+ 前端图片懒加载
+ 登录，注册
+ 后台文章管理，用户管理
+ 图片上传

## Asy 1.3 Released 2019年03月24日
+ 改用iris框架
+ cookie登录加密
+ 文件目录结构调整
+ 管理模块引入 webpack + vue + vue-router + axios

## 安装说明
1. 安装mysql
2. 使用iissy.com.sql初始化数据库以及数据
3. 修改数据库连接（src/iissy.com/utils/config.go）

## 环境变量设置
如果你将代码下载到了D:\github.com\goweb，添加D:\github.com\goweb到GOPATH环境变量，如果有多个GOPATH目录，请确保最后没有分号

## 组件下载
1. go get github.com/go-sql-driver/mysql
2. go get github.com/kataras/iris
3. go get github.com/gorilla/securecookie
4. 运行命令启动# go build -o main.exe ./ && main.exe

## 使用说明
+ 请确保80端口没有被别的服务（比如iis）占用，然后在浏览器中输入：http://localhost
+ 登陆账号
+ 账号：jimmy 密码：123456

## 首页
![markdown](https://github.com/iissy/goweb/blob/master/public/home.png "首页图片")

## 文章
![markdown](https://github.com/iissy/goweb/blob/master/public/art.png "文章图片")

## 发布
![markdown](https://github.com/iissy/goweb/blob/master/public/add.png "发布图片")

## 列表
![markdown](https://github.com/iissy/goweb/blob/master/public/list.png "列表图片")

2019年5月下旬

> 将发布用户分角色权限管理功能

## 相关网站
+ 程序员网址导航：https://www.hrefs.cn
