---
title: go modulesのファイル分割とMonorepoについて
category: "monorepo"
cover: golangimage.png
author: Takanori Fukuyama
---

2020/08/02 更新

go modulesを検証した

go modulesのinitialize

```fish
$ go mod init micrioservices
```

とすると

go.modファイルが作られる
```mod
module microservices

go 1.14
```


### ディレクトリ構造

```shell
project_dir
├── core_domains
├── generic_subdomains
│   ├── accounting_service
│   │   ├── README.md
│   │   ├── docker-compose.yml
│   │   ├── domain
│   │   ├── framework_driver
│   │   ├── interface_adapter
│   │   ├── k8s.yml
│   │   ├── main.go
│   │   ├── skaffold.yml
│   │   └── usecase
│   └── authority_service
│       ├── domain
│       └── usecase
│       ├── interface_adapter
│       ├── framework_driver
│       ├── BUILD.bazel
│       └── main.go
├── go.mod
├── go.sum
└── vendor
```


## ファイルのimport

例えば

`microservices/generic_subdomains/accounting_service/main.go`
```go
package main

import (
    "microservices/generic_subdomains/accounting_service/interface_adapter/gateway/user_repository"
)

func main() {
  gateway.NewUserRepositoryImol( ... )
  ...
}

```


![golang](./golangimage.png)
