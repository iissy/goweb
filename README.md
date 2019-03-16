## 安装说明
1. 安装mysql
2. 使用iissy.com.sql初始化数据库以及数据
3. 修改数据库连接（src/iissy.com/utils/config.go）

## 环境变量设置
如果你将代码下载到了D:\github.com\goweb，添加D:\github.com\goweb到GOPATH环境变量，如果有多个GOPATH目录，请确保最后没有分号

## 组件下载
1. go get github.com/go-sql-driver/mysql
2. go get github.com/julienschmidt/httprouter
3. 运行命令启动# go run main.go

## 使用说明
+ 请确保80端口没有被别的服务（比如iis）占用，然后在浏览器中输入：http://localhost
+ 登陆账号
+ 账号：jimmy 密码：123456

## 首页
![markdown](https://github.com/iissy/goweb/blob/master/public/home.png "首页图片")

## 发布
![markdown](https://github.com/iissy/goweb/blob/master/public/add.png "发布图片")

## 列表
![markdown](https://github.com/iissy/goweb/blob/master/public/list.png "列表图片")


## 程序员网址导航：https://www.hrefs.cn

# 将在几天内发布一个更新的版本，使用iris web框架改造过的，代码基本完成

# 将在一个月内将内容管理部分改造成支持webpack + vue + vue-router + axios单页面
