---
title: go-study
date: 2022/07/09 18:25:16
categories:
  - 成长经历
tags:
  - golang
---

## 安装

win10：[详见](https://golang.google.cn/dl/)

等待安装完成后，cmd 输入`go version`，有版本输出，则验证安装成功。

> Tips
>
> Go 安装过程中，会自动配置好环境变量

### Goproxy 配置

配置 Go 模块代理，访问下载更快。

```sh
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.cn,direct
```

[其他 Go 模块代理平台](https://goproxy.io/zh/)

## IDE

首推 VSCode，因为开源免费，功能强大。（微软出品，必属精品）

**VSCode 插件**：Go

**VSCode 设置**：

编辑器

@lang:go Insert Spaces

- [x] 按 Tab 键时插入空格

Go 插件

```sh
# Format Tool
选择 goformat
```

首次写入`.go`文件，引入的依赖库，如需访问，根据 VSCode 的提示下载对应的 Go 插件。

## 依赖管理

### 三个阶段

#### GOPATH

- GO111MODULE=off

- 默认在`~/go`(unix,linux)，`%USERPROFILE%\go`(windows)

- 必须都放在`src`目录下。

怎么解决版本问题？

在项目中创建`vendor`目录，就近原则查找依赖。

#### GOVENDOR

- 每个项目有自己的`vendor`目录，存放第三方库
- 大量第三方依赖管理工具：[glide](https://github.com/Masterminds/glide)，dep，go dep，...

#### GOMOD(推荐)

前面2个逐渐被淘汰。

由go命令统一管理，用户不必关心目录结构

- 初始化：go mod init
- 增加依赖：go get
- 更新依赖：go get [@v...], go mod tidy
- 将旧项目迁移到go mod：go mod init, go build ./...

## 目录整理

一个包只能又一个main入口函数

> go-base目录下的文件作为demo使用（不规范的），实际使用必须一个包只能有一个main入口函数。

## 参考

[1] [Golang Documentation](https://golang.google.cn/doc/)

[2] [Goproxy.cn](https://goproxy.cn/)

[3] [李文周的博客](https://www.liwenzhou.com/posts/Go/golang-menu/)

[4] [8 小时转职 Golang 工程师(如果你想低成本学习 Go 语言](https://www.bilibili.com/video/BV1gf4y1r79E?from=search&seid=14989564876573827402&spm_id_from=333.337.0.0)

[5] [Go 语言文档](https://www.topgoer.com/)

[6] [一文搞懂Go语言的最新依赖管理：go mod的使用](https://blog.csdn.net/Sihang_Xie/article/details/124851399)
