# 抖音开放平台API (In progress)

## Docs 文档

[![Go Report Card](https://goreportcard.com/badge/github.com/amorist/douyin)](https://goreportcard.com/report/github.com/amorist/douyin)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/amorist/douyin?tab=doc)

## Usage 示例

```go
scopes := oauth.GetAllScope(), 
dy := douyin.NewDouyin(&config.Config{
    ClientKey:    "your client key",
    ClientSecret: "your client secret",
    RedirectURL:  "your redirect url",
    Scopes:  scopes,
    Cache:   util.NewMemCache(),
})

oauth := dy.GetOauth()
url := oauth.GetRedirectURL("amorist")
```

## Issue 如有问题，可以提issue或通过微信联系我

加微信`amor-ist`备注`douyin api`.
