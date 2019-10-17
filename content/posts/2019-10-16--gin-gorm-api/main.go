package main

import (
    "os"
    
    "github.com/gin-gonic/gin"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"

    "github.com/google/uuid"
)

func main() {
	dbMigrate(connection())
	defer connection().Close()

    r := newRouter()
    r.Run(":8080")
}

func newRouter() *gin.Engine {
    r := gin.Default()

    r.POST("/v1/core/user", func(c *gin.Context) {
        uuidObj, _ := uuid.NewUUID()
        UserID := uuidObj.String()

        var req tUserAuth
        c.BindJSON(&req)

        auth := &tUserAuth{
            UserID: UserID,
            Line: req.Line,
            Google: req.Google,
            Facebook: req.Facebook,
            Profile: tUser{
                UserID: UserID,
                Name: req.Profile.Name,
            },
        }
        connection().Create(auth)
        var u tUserAuth
        query := connection().Last(&u)
        defer connection().Close()
        c.JSON(200, query)
    })

    r.GET("/v1/core/user", func(c *gin.Context) {
        result := connection().Find(&tUserAuth{})
        c.JSON(200, result)
        connection().Close()
    })

    r.GET("/v1/core/user/:id", func(c *gin.Context) {
        id := c.Param("id")
        result := connection().Where("aikasa = ?", id).First(&tUserAuth{})
        c.JSON(200, result)
        connection().Close()
    })

    r.DELETE("/v1/core/user/:id", func(c *gin.Context) {
        id := c.Param("id")
        connection().Where("aikasa = ?", id).Delete(tUserAuth{})
        connection().Close()
        c.JSON(200, gin.H{
            "status": "OK",
        })
    })

    return r
}
func dbMigrate(db *gorm.DB) *gorm.DB {
    db.AutoMigrate(&tUserAuth{}, &tUser{})
    return db
}

type tUser struct {
    gorm.Model
    UserID string `json:"user_id"`
    Name   string `json:"name"`
}

type tUserAuth struct {
    gorm.Model
    UserID    string `json:"user_id"`
    Line      string `json:"line_id"`
    Facebook  string `json:"facebook_id"`
    Google    string `json:"google_id"`
    Profile   tUser `json:"profile"`
}

func connection() *gorm.DB{
    DBMS := "mysql"
    USER := "root"
    PASS := os.Getenv("MYSQL_PASSWORD")
    PROTOCOL := "tcp(mysql:3306)"
    DBNAME := "aikasa_db"

    CONNECT := USER + ":" + PASS + "@" + PROTOCOL + "/" + DBNAME + "?parseTime=true"
    db, err := gorm.Open(DBMS, CONNECT)
    if err != nil {
        panic(err.Error())
    }
    return db
}