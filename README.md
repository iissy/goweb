# 转移仓库
```
https://github.com/go-water/go-water
```
<div style="color:#ff0000">升级的技术框架，已经转移到以上仓库。</div>

### 简介
一个由Go语言写的程序员网址导航，聚合go生态比较热门的技术栈。

## 星星增长趋势
[![Stargazers over time](https://starchart.cc/iissy/goweb.svg)](https://starchart.cc/iissy/goweb)

#### asy 2.1.2 Released 2021年02月11日
+ 仓库更新到最新的 go-micro
+ 将微服务改成 gRPC
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
+ ORM 框架引入 gorp
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

## 使用说明
+ 请确保 80 端口没有被别的服务（比如iis）占用，然后在浏览器中输入：http://localhost
+ 如果需要添加修改数据，必须搭建内容管理项目，见安装说明

## 官方产品
+ https://jitask.com
+ https://go-water.cn
