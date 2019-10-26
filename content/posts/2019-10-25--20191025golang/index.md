---
title: go modulesのファイル分割について
category: "golang"
cover: golangimage.png
author: Takanori Fukuyama
---

![golang](./golangimage.png)

go modulesでファイル分割をしようと思ったけど、なかなか難しかった。
色々検証してみた。

```fish
$ go env
```
を実行すると

### 最終的なディレクトリ構造




```
project
├──  database
│  └──  database.go
├──  docker-compose.yml
├──  Dockerfile
├──  go.mod
├──  go.sum
├──  main.go
├──  models
│  └──  models.go
├──  README.md
├──  run.sh
└──  project
```

