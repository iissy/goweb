这是一个由Go语言写的，简单的CMS（内容管理系统），以应用到爱斯园（程序员网址导航）。

## 星星增长趋势
[![Stargazers over time](https://starchart.cc/iissy/goweb.svg)](https://starchart.cc/iissy/goweb)

#### asy 2.0 Released 2020年02月23日
+ 技术大更新，引入流行的 web 端框架
+ 源码应用到——爱斯园（程序员网址导航）
+ 前后端分离，内容管理项目移入另一个仓库(https://github.com/iissy/hrefs.cn)
+ 内容管理项目用到的接口在此项目，它仅仅是一个 vue 项目
+ 此项目可以独立跑，而后台项目必须依赖次项目接口

### 技术栈
+ go iris 框架
+ cookie 认证，双重加密
+ 前后端分离
+ 分组路由
+ 中间件实现授权访问
+ redis 管理登陆会话
+ 日志错误管理
+ ORM 框架引入
+ 图片上传，uuid命名
+ 使用 Go Module 方式开发

### 功能点
+ 原创文章
+ 优秀网摘
+ 常用网址

## 安装说明
1. 安装 mysql
2. 使用hrefs.cn.sql初始化数据库以及数据
3. 安装 redis（不启用后台管理可以不装）
4. 修改数据库连接（conf/config.json）
5. 后台代码仓库：https://github.com/iissy/hrefs.cn

## 使用代码
由于大陆网络无法下载google的包，使用七牛公司的代理，下载包前执行命令
1. go env -w GOPROXY=https://goproxy.cn,direct

## 组件下载
1. go get github.com/go-sql-driver/mysql
2. go get github.com/kataras/iris
3. go get github.com/gorilla/securecookie
4. go get github.com/go-gorp/gorp
5. go get github.com/gomodule/redigo
6. go get github.com/google/uuid
7. go get github.com/juju/errors
8. go get github.com/kataras/golog

#### 运行命令启动> go run main.go

## 使用说明
+ 请确保80端口没有被别的服务（比如iis）占用，然后在浏览器中输入：http://localhost
+ 如果需要添加修改数据，必须使用后台代码，见安装说明

## 打赏站长
![markdown](https://www.hrefs.cn/payme.jpg)

## 官方网站
+ 程序员网址导航：https://www.hrefs.cn