这是一个由Go语言写的，简单的内容管理系统，包括前端显示，支持SEO，图片懒加载，后台文章管理，图片上传等功能的网站，也是一个供Go语言的学习者参考的一个demo。版本仍在继续更新维护，见下面版本发布时间。

## Asy 1.2 Released 2019年03月17日
+ 改用iris框架
+ cookie登录加密
+ 文件目录结构调整

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

2019年4月7日

> 将内容管理部分改造成支持 webpack + vue + vue-router + axios

## 程序员网址导航：https://www.hrefs.cn

我是一个全职社区贡献者，感谢您的打赏！
![markdown](https://github.com/iissy/goweb/blob/master/public/pay.jpg "向我支付")
