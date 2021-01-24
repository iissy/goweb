一个由Go语言写的程序员网址导航，由于国内政策不允许个人备案导航相关的网站，目前网站放在海外的阿里云服务器上，所以访问速度稍微有点慢。

## 出售 https://www.hrefs.cn
```
价格：2000元
联系QQ：17660121
由于本人忙碌，网站最近少更新，低价转让网站，所有资源，包括源码，阿里云海外服务器（1核心1G内存），域名一并转让。
```

## 星星增长趋势
[![Stargazers over time](https://starchart.cc/iissy/goweb.svg)](https://starchart.cc/iissy/goweb)

#### asy 2.0.2 Released 2020年04月05日
+ 完成 iris、gin、go-micro 技术大升级
+ iris 显示内容渲染，gin 实现接口，go-micro 微服务
+ 一个应用启用三个端口，web、api、srv各占一个
+ 内容管理项目移入另一个独立仓库
+ 内容管理仓库(https://github.com/iissy/hrefs.cn)
+ 内容管理项目用到的接口在此项目
+ 内容展示这个项目可以独立跑
+ 旧代码见 old 分支

### 技术栈
+ go iris 网页 / go gin 接口
+ go-micro （内置gRPC）微服务框架
+ consul 服务注册发现
+ 分组路由
+ 中间件实现授权访问
+ redis 管理登陆会话
+ 日志错误管理
+ ORM 框架 gorp 引入
+ 图片上传，google.uuid 生成文件名
+ 使用 Go Module 方式开发

### 内容包含
+ 技术原创文章
+ 优秀技术网摘
+ 程序员常用网址

## 安装说明
1. 安装 mysql
2. 使用 hrefs.cn.sql 初始化数据库以及数据
3. 安装 redis（不启用内容管理可以不装）
4. 修改配置（conf/config.json）
5. 内容管理代码仓库：https://github.com/iissy/hrefs.cn

## 组件下载配置
由于大陆网络无法下载google的包，使用七牛公司的代理，下载包前执行命令
1. go env -w GOPROXY=https://goproxy.cn,direct

## 组件下载
```
go get github.com/go-sql-driver/mysql
go get github.com/kataras/iris
go get github.com/go-gorp/gorp/v3
go get github.com/gomodule/redigo
go get github.com/google/uuid
go get github.com/juju/errors
go get github.com/kataras/golog
go get github.com/micro/go-micro/v2
go get github.com/micro/go-micro/v2/config
go get github.com/micro/go-micro/v2/registry
go get github.com/micro/go-plugins/registry/consul/v2
```

## 使用说明
+ 请确保 80 端口没有被别的服务（比如iis）占用，然后在浏览器中输入：http://localhost
+ 如果需要添加修改数据，必须搭建内容管理项目，见安装说明

## 官方网站
+ 程序员网址导航：https://www.hrefs.cn