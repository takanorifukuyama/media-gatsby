---
title: golangのAPI作った(Gin,Gorm)
category: "golang"
cover: photo-1465070845512-2b2dbdc6df66.jpg
author: Takanori Fukuyama
---

### ソースコード

```go
package main

import (
    "os"
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "github.com/google/uuid"
)

func main() {
    r := NewRouter()
    r.Run(":8080")
}
```

とりあえずテスト