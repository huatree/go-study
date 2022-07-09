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

## 参考

[1] [Golang Documentation](https://golang.google.cn/doc/)

[2] [Goproxy.cn](https://goproxy.cn/)

[3] [李文周的博客](https://www.liwenzhou.com/posts/Go/golang-menu/)

[4] [8 小时转职 Golang 工程师(如果你想低成本学习 Go 语言](https://www.bilibili.com/video/BV1gf4y1r79E?from=search&seid=14989564876573827402&spm_id_from=333.337.0.0)

[5] [Go 语言文档](https://www.topgoer.com/)
